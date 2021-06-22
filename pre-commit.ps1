$NAME = 'pre-commit-docker'
docker run -t --rm --name $NAME -v ${PWD}:/pre-commit taghash/pre-commit:latest

# -v $HOME/.aws:/root/.aws:ro
