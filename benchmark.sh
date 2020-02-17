go build
cd examples
echo "dir: " $(pwd)
pandoc -v
echo "pandoc beamer:"
hyperfine --show-output "pandoc -i example.md -o out.pdf"
echo "point:"
hyperfine --show-output "../point -i example.md -o out.html"
