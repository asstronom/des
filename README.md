# des

to build project run: go build -o encrypt.exe

# flags

-decrypt      decrypts your file
-triple       uses TDES as a cipher
-i            specifies path to input file
-o            specifies path to output file
-key          specifies key

# examples of usage

usage (DES):
./encrypt.exe -key asimovak -i sample.txt
./encrypt.exe -key asimovak -i output.txt -o output_dec.txt -decrypt 

usage (TDES):
./encrypt.exe -key asimovakpudgeboosterenko -triple -o output_dec_triple.txt -i output_triple.txt
./encrypt.exe -key asimovakpudgeboosterenko -triple -o output_dec_triple.txt -i output_triple.txt -decrypt 
