package main

import (
	"os"
	"path/filepath"

	"github.com/goyek/goyek"
)

// BuildRoot reflects the directory above `build` which should be the project directory. This variable provides more predictable path handling once set for all subsequent tasks.
var BuildRoot string

// GetBuildRoot navigates up from `build` directory to ensure the path for the project is globally available for tasks with a simple call.
func GetBuildRoot(tf *goyek.TF) {
	wd, err := os.Getwd()
	if err != nil {
		tf.Errorf("getwd: [%v]", err)
	}
	// WITH HELPER: BuildRoot = resolveParentDirectory(wd)
	// projectDirectory := filepath.Join("../", wd)
	projectDirectory := resolveParentDirectory(tf, wd)
	BuildRoot = resolveABSPath(tf, projectDirectory)
	// BuildRoot, err := filepath.Abs(projectDirectory)
	// if err != nil {
	// 	tf.Errorf("filepath.Abs(ProjectDirectory): [%v]", err)
	// }
	tf.Logf("BuildRoot: [%s]", BuildRoot)
}

// ResolveParentDirectory returns the directory above the provided directory as a fully qualified absolute path
func resolveParentDirectory(tf *goyek.TF, childDirectory string) (parentDirectory string) {
	projectDirectory := filepath.Dir(childDirectory)
	parentDirectory, err := filepath.Abs(projectDirectory)
	if err != nil {
		tf.Errorf("filepath.Abs(ProjectDirectory): [%v]", err)
	}
	tf.Logf("childDirectory [%s] --> parentDirectory: [%s]", childDirectory, parentDirectory)
	return parentDirectory
}

// resolveABSPath returns absolute path of any path, and logs error upon failure
func resolveABSPath(tf *goyek.TF, directory string) (ABSPath string) {
	ABSPath, err := filepath.Abs(directory)
	if err != nil {
		tf.Errorf("ABSPath: [%v]", err)
	}
	tf.Logf("directory [%s] --> ABSPath: [%s]", directory, ABSPath)
	return ABSPath
}
