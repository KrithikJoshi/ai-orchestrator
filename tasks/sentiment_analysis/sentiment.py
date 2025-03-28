from textblob import TextBlob

input_path = '/data/output.txt'
output_path = '/data/sentiment.txt'

with open(input_path, 'r') as f:
    text = f.read()

blob = TextBlob(text)
sentiment = blob.sentiment.polarity

if sentiment > 0:
    result = "Positive"
elif sentiment < 0:
    result = "Negative"
else:
    result = "Neutral"

with open(output_path, 'w') as f:
    f.write(f"Sentiment: {result} (score: {sentiment})")
