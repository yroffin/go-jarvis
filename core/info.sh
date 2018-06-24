SHA1=$(git log -g --pretty=oneline -v -1 | cut -f1 -d' ')
echo {
echo "    "\"build-date\": \"`date`\",
echo "    "\"sha-1\": \"$SHA1\"    
echo }