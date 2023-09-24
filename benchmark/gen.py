# this file is used to generate the benchmark data

import json
import os, datetime, sys, urllib.parse
import random, string

host = "http://localhost:8080/"
shortern = host + "shortern"

class json_generator:
    def __init__(self, url:str, path:str, timestamp:datetime.datetime):
        self.filename = path + ".json"
        self.data = {
            "URL": url,
            "Expiration": timestamp.isoformat(),
        }
    def writeJSON(self):
        data = json.dumps(self.data)
        with data_writer(self.filename) as f:
            f.write(data)

def get_post_writer():
    return open("benchmark/post.txt", "a")
def get_get_writer():
    return open("benchmark/get.txt", "a")

def data_writer(filename:str):
    path = "benchmark/data"
    # Check whether the specified path exists or not
    isExist = os.path.exists(path)
    #printing if the path exists or not
    if not isExist:
        # Create a new directory because it does not exist
        os.makedirs(path)
    return open("{}/{}".format(path, filename), "a")   

def get_url()->(str, str):
    base = "http://localhost:8080"
    path = ''.join(random.choice(string.ascii_letters+string.digits) for x in range(10))
    res = urllib.parse.urljoin(base, path)
    return res, path
    

if __name__ == "__main__":
    count = int(sys.argv[1])
    today = datetime.datetime.now(datetime.timezone.utc)
    tmr = today + datetime.timedelta(days=1)
    yesterday = today - datetime.timedelta(days=1)
    with get_post_writer() as f:
        with get_get_writer() as g:
            # generate post files and valid get files
            for i in range(count):
                url, path = get_url()
                a = json_generator(url, path, random.choice([today, tmr, yesterday]))
                a.writeJSON()
                f.writelines(["POST {}\n".format(shortern), "@benchmark/data/{}\n".format(a.filename), "\n"])
                g.writelines(["GET {}\n".format(url), "\n"])
            # generate invalid get files
            for i in range(count*30):
                url, _ = get_url()
                g.writelines(["GET {}\n".format(url), "\n"])
