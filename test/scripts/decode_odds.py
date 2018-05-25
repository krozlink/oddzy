import json

obj = json.loads(editor.getText())

response = obj["r"]
decoded = response.encode('rot13').decode('base64')

jsonResult = json.loads(decoded)
stringResult = json.dumps(jsonResult, indent=2)

editor.setText(stringResult)


# Requires Python Script extension