for i in {user,video,favorite,comment,relation}
do
  cd output/$i/output
  mkdir log
  nohup nohup sh bootstrap.sh > log/runtime.log 2>&1 &
  cd -
done

cd output/api
mkdir log
nohup ./main > log/runtime.log 2>&1 &
