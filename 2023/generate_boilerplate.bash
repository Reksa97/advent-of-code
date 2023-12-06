# give executable permissions with chmod +x generate_boilerplate.bash
if [ $# -eq 0 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

day=$1

mkdir -p day$day
mkdir -p inputs

cp day0/main.go day$day/main.go
sed -i '' "s/day := 0/day := $day/g" day$day/main.go

touch test_inputs/$day.input
touch inputs/$day.input