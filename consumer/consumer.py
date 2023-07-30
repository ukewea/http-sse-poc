import requests

def event_stream():
    url = "http://localhost:8080/events"
    try:
        response = requests.get(url, stream=True, timeout=5)
        if response.status_code == 200:
            for line in response.iter_lines():
                if line:
                    decoded_line = line.decode('utf-8')
                    print(decoded_line)

            print("event terminated")
        else:
            print("error: status code {}".format(response.status_code))
    except requests.exceptions.Timeout:
        print("error: request timed out")

if __name__ == "__main__":
    event_stream()