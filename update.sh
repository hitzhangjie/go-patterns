#!/bin/bash 

if [ ! -f go.work ]
then
    echo "init go workspace"
    go work init
fi

echo "update go.work"

for i in `find -iname go.mod`
do 
    echo ${i%%/go.mod}; 
    go work use "${i%%/go.mod}"; 
done

