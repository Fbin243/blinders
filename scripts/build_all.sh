
if [[ "$1" != "dev" && "$1" != "staging" && "$1" != "prod" ]]; 
then
    echo "Usage: $0 with one of 'dev|staging|prod'"
    exit 1
fi

sh scripts/build_golambda.sh $1