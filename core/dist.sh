rm -rf dist/*
mkdir -p dist/jarvis-ui
cp ../ui/dist.tar dist/dist.tar
cp -rf ../swagger-ui dist
cd dist/jarvis-ui && tar xvf ../dist.tar && rm -f ../dist.tar
