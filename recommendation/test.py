import requests

for i in range(100):
    response = requests.get("http://127.0.0.1:5000/quarks", json = {'UserID' : 7})
    print(response)
    print(response.json())