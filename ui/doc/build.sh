# postman collection 2.1

if (( "$#" != 1 ))
then
    echo "to name of doc"

    exit 1
fi

docgen build --in ./$1 --out ./$1.html