# required:
# env starts with KUBE_LABEL_

unset IFS

echo > patch.json
echo "[" >> patch.json
for var in $(compgen -e | grep KUBE_LABEL_); do
  echo "$var = ${!var}"
  label_name="${var//KUBE_LABEL_/}"
  {
    echo '{';
    echo '  "op":"add", "path":"/metadata/labels/'"$label_name"'", "value":"'"${!var}"'"'
    echo '},'
  } >> patch.json
done
sed -i '$ s/.$//' patch.json
echo "]" >> patch.json
cat patch.json
