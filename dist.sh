LATEST=$(curl -s https://github.com/yroffin/ng-jarvis/releases/latest -s | cut -f2 -d\" | sed s:/tag/:/download/:)/dist.tar
echo $LATEST
rm -rf dist/*
mkdir -p dist/jarvis-ui
curl -s -L $LATEST -o dist/dist.tar
cp -rf resources/static/swagger-ui dist
cd dist/jarvis-ui && tar xvf ../dist.tar && rm -f ../dist.tar
