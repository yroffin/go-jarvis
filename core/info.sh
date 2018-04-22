SHA1=$(git log -g --pretty=oneline -v -1 | cut -f1 -d' ') && echo $SHA1
echo {
echo "    "\"version\": \"$SHA1\"    
echo }