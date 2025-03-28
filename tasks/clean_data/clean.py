import re
import sys

input_path = '/data/input.txt'
output_path = '/data/output.txt'

with open(input_path, 'r') as f:
    text = f.read()

# Clean: lower, remove special chars
cleaned = re.sub(r'[^a-zA-Z0-9\s]', '', text.lower())

with open(output_path, 'w') as f:
    f.write(cleaned)
