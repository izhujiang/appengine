fd --extension go --extension mod | xargs sed -i "" -e \
    's/github.com\/izhujiang\/appengine/github.com\/izhujiang\/myapp/g'
