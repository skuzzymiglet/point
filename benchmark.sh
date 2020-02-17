go build
cd examples
echo "dir: " $(pwd)
pandoc -v
echo "pandoc beamer:"
hyperfine "pandoc -i example.md -o out.pdf"
echo "point:"
hyperfine "../point -i example.md -o out.html"
