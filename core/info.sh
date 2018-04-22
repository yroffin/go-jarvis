SHA1=$(git log -g --pretty=oneline -v -1 | cut -f1 -d' ')
echo {
echo "    "\"version\": \"$SHA1\"    
echo }