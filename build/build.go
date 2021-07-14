package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/goyek/goyek"
)

// var (
// // ArtifactDirectory is the build output directory for all binaries and other files. This simplifies project management instead of having possible binaries in each directory go files are built from.
// // ArtifactDirectory = "artifacts"

// )

// buildDir is the current build directory for the goyek files.
const buildDir = "build"

// ArtifactDirectory is the build output directory for all binaries and other files. This simplifies project management instead of having possible binaries in each directory go files are built from.
const ArtifactDirectory = "artifacts"

// dockerFilePrecommit is the pre-commit tooling dockerfile to use.
const dockerFilePrecommit = "Dockerfile.v2.precommit"

// BuildRoot is the absolute path for the project directory, removing the need to figure out relative path starting points.
// const BuildRoot = "../"

func main() {
	flow().Main()
}

func flow() *goyek.Taskflow {
	flow := &goyek.Taskflow{}

	// parameters
	ci := flow.RegisterBoolParam(goyek.BoolParam{
		Name:  "ci",
		Usage: "Whether CI is calling the build script",
	})

	// tasks
	clean := flow.Register(taskClean())
	build := flow.Register(taskBuild())
	fmt := flow.Register(taskFmt())
	markdownlint := flow.Register(taskMarkdownLint())
	misspell := flow.Register(taskMisspell())
	golangciLint := flow.Register(taskGolangciLint())
	test := flow.Register(taskTest())
	modTidy := flow.Register(taskModTidy())
	diff := flow.Register(taskDiff(ci))

	precommitRunAll := flow.Register(taskPrecommitRunAll())

	_ = flow.Register(taskDockerBuild())
	// pipelines
	lint := flow.Register(taskLint(goyek.Deps{
		misspell,
		markdownlint,
		golangciLint,
		precommitRunAll,
	}))
	_ = flow.Register(taskAll(goyek.Deps{
		clean,
		build,
		fmt,
		lint,
		test,
		modTidy,
		diff,
	}))
	// flow.DefaultTask = all

	return flow
}

func taskClean() goyek.Task {
	return goyek.Task{
		Name:    "clean",
		Usage:   "remove git ignored files",
		Command: goyek.Exec("git", "clean", "-fX"),
	}
}

func taskBuild() goyek.Task {
	return goyek.Task{
		Name:    "build",
		Usage:   "go build",
		Command: goyek.Exec("go", "build", "./..."),
	}
}

func taskFmt() goyek.Task {
	return goyek.Task{
		Name:  "fmt",
		Usage: "gofumports",
		Command: func(tf *goyek.TF) {
			installFmt := tf.Cmd("go", "install", "mvdan.cc/gofumpt/gofumports")
			installFmt.Dir = buildDir
			if err := installFmt.Run(); err != nil {
				tf.Fatalf("go install gofumports: %v", err)
			}
			tf.Cmd("gofumports", strings.Split("-l -w -local github.com/goyek/goyek .", " ")...).Run()
		},
	}
}

// taskDockerBuild runs docker build commands against all local dockerfiles.
func taskDockerBuild() goyek.Task {
	return goyek.Task{
		Name:  "dockerbuild",
		Usage: "pull any docker dependencies in project",

		Command: func(tf *goyek.TF) {
			GetBuildRoot(tf)
			buildPrecommit := tf.Cmd("docker", "build", "--pull", "--rm", "-f", dockerFilePrecommit, "-t", "precommit-custom:latest", ".")
			tf.Logf("ℹ️ buildPrecommit > %s", buildPrecommit.String())
			// docker build --pull --rm -f "Dockerfile.precommit" -t precommit-custom:latest "."

			buildPrecommit.Dir = BuildRoot
			tf.Logf("ℹ️ buildPrecommit dir set to: [%s]\n", BuildRoot)

			if err := buildPrecommit.Run(); err != nil {
				tf.Fatalf("❗ docker build Dockerfile.precommit [%v]", err)
			}
			tf.Log("✅ docker precommit image built")
			// tf.Cmd("gofumports", strings.Split("-l -w -local github.com/goyek/goyek .", " ")...).Run() // nolint // it is OK if it returns error
		},
	}
}

// taskPrecommitRunAll runs docker precommit image against all files in repo.
func taskPrecommitRunAll() goyek.Task {
	return goyek.Task{
		Name:  "precommit-runall",
		Usage: "runs docker precommit image against all files in repo",

		Command: func(tf *goyek.TF) {
			GetBuildRoot(tf)
			runPrecommit := tf.Cmd("docker", "run", "--rm", "-v", BuildRoot+":/pre-commit", "precommit-custom:latest")
			tf.Logf("ℹ️ taskPrecommitRunAll > %s", runPrecommit.String())

			// docker build --pull --rm -f "Dockerfile.precommit" -t precommit-custom:latest "."

			runPrecommit.Dir = BuildRoot
			tf.Logf("ℹ️ taskPrecommitRunAll dir set to: [%s]\n", BuildRoot)

			if err := runPrecommit.Run(); err != nil {
				tf.Fatalf("❗ docker run taskPrecommitRunAll [%v]", err)
			}
			tf.Log("✅ taskPrecommitRunAll")
			// tf.Cmd("gofumports", strings.Split("-l -w -local github.com/goyek/goyek .", " ")...).Run() // nolint // it is OK if it returns error
		},
	}
}

// func TaskGoTooling() goyek.Task {
// 	return goyek.Task{
// 		Name:  "gotools",
// 		Usage: "ensure all go tooling dependencies are installed for development work",
// 		Command: func(tf *goyek.TF) {
// 			tf.Log("Installing go tooling for development")
// 			for _, i := range goToolsRepos {
// 				p.Title = "Installing " + i
// 				GetCommand := tf.Cmd("go", "install", i)
// 				if err := GetCommand.Run(); err != nil {
// 					pterm.Warning.Printf("Could not install [%s] per [%v]\n", i, err)
// 				}
// 				p.Increment()
// 				time.Sleep(second / 2)
// 			}
// 			p.Title = "gotools successfully installed"
// 			_, _ = p.Stop()
// 		},
// 	}
// }

func taskMarkdownLint() goyek.Task {
	return goyek.Task{
		Name:  "markdownlint",
		Usage: "markdownlint-cli (requires docker)",
		Command: func(tf *goyek.TF) {
			curDir, err := os.Getwd()
			if err != nil {
				tf.Fatal(err)
			}

			docsMount := curDir + ":/markdown"
			if err := tf.Cmd("docker", "run", "-v", docsMount, "06kellyjac/markdownlint-cli:0.27.1", "**/*.md").Run(); err != nil {
				tf.Error(err)
			}

			gitHubTemplatesMount := filepath.Join(curDir, ".github") + ":/markdown"
			if err := tf.Cmd("docker", "run", "-v", gitHubTemplatesMount, "06kellyjac/markdownlint-cli:0.27.1", "**/*.md").Run(); err != nil {
				tf.Error(err)
			}
		},
	}
}

func taskMisspell() goyek.Task {
	return goyek.Task{
		Name:  "misspell",
		Usage: "misspell",
		Command: func(tf *goyek.TF) {
			installFmt := tf.Cmd("go", "install", "github.com/client9/misspell/cmd/misspell")
			installFmt.Dir = buildDir
			if err := installFmt.Run(); err != nil {
				tf.Fatalf("go install misspell: %v", err)
			}
			lint := tf.Cmd("misspell", "-error", "-locale=US", "-i=imports", ".")
			if err := lint.Run(); err != nil {
				tf.Fatalf("misspell: %v", err)
			}
		},
	}
}

func taskGolangciLint() goyek.Task {
	return goyek.Task{
		Name:  "golangci-lint",
		Usage: "golangci-lint",
		Command: func(tf *goyek.TF) {
			installLint := tf.Cmd("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint")
			installLint.Dir = buildDir
			if err := installLint.Run(); err != nil {
				tf.Fatalf("go install golangci-lint: %v", err)
			}
			lint := tf.Cmd("golangci-lint", "run")
			if err := lint.Run(); err != nil {
				tf.Fatalf("golangci-lint run: %v", err)
			}
		},
	}
}

func taskTest() goyek.Task {
	return goyek.Task{
		Name:    "test",
		Usage:   "go test with race detector and code covarage",
		Command: goyek.Exec("go", "test", "-race", "-covermode=atomic", "-coverprofile=coverage.out", "./..."),
	}
}

func taskModTidy() goyek.Task {
	return goyek.Task{
		Name:  "mod-tidy",
		Usage: "go mod tidy",
		Command: func(tf *goyek.TF) {
			if err := tf.Cmd("go", "mod", "tidy").Run(); err != nil {
				tf.Errorf("go mod tidy: %v", err)
			}

			toolsModTidy := tf.Cmd("go", "mod", "tidy")
			toolsModTidy.Dir = buildDir
			if err := toolsModTidy.Run(); err != nil {
				tf.Errorf("go mod tidy: %v", err)
			}
		},
	}
}

func taskDiff(ci goyek.RegisteredBoolParam) goyek.Task {
	return goyek.Task{
		Name:   "diff",
		Usage:  "git diff",
		Params: goyek.Params{ci},
		Command: func(tf *goyek.TF) {
			if !ci.Get(tf) {
				tf.Skip("ci param is not set, skipping")
			}

			if err := tf.Cmd("git", "diff", "--exit-code").Run(); err != nil {
				tf.Errorf("git diff: %v", err)
			}

			cmd := tf.Cmd("git", "status", "--porcelain")
			sb := &strings.Builder{}
			cmd.Stdout = io.MultiWriter(tf.Output(), sb)
			if err := cmd.Run(); err != nil {
				tf.Errorf("git status --porcelain: %v", err)
			}
			if sb.Len() > 0 {
				tf.Error("git status --porcelain returned output")
			}
		},
	}
}

func taskLint(deps goyek.Deps) goyek.Task {
	return goyek.Task{
		Name:  "lint",
		Usage: "all linters",
		Deps:  deps,
	}
}

func taskAll(deps goyek.Deps) goyek.Task {
	return goyek.Task{
		Name:  "all",
		Usage: "build pipeline",
		Deps:  deps,
	}
}
