git ls-files --full-name -- "**/go.mod" | xargs -n 1 dirname | xargs -I{} bash -c 'cd "{}" && go mod tidy'