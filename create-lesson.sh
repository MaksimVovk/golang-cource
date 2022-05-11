if [ "$(git diff)" ]; then
  echo 'there are local modifications';
  exit -1;
fi

if [ "$(git branch | grep \* | cut -d ' ' -f2)" != "master" ]; then
  echo 'current branch is not master';
  git checkout master;
fi

git pull
git checkout -b $block_name

(cd ./src/lectures; mkdir ./demo/$block_name; touch ./demo/$block_name/$block_name.go)
(cd ./src/lectures; mkdir ./exercise/$block_name; touch ./exercise/$block_name/$block_name.go)

echo 'created ./demo/$block_name/$block_name.go';
echo 'created ./exercise/$block_name/$block_name.go';
