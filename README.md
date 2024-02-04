# YTT Ext

### simple patching without forking the entire ytt repo

```sh
# change this to desired ytt tag
export TAG=v0.0.0

# just for reference
echo "https://github.com/carvel-dev/ytt/blob/$TAG"

# test locally
goreleaser release --clean --snapshot

# push to ci
git tag -f "$TAG"
git push --tags
```
