from flask import Flask
  
app = Flask(__name__)

def hello_world():
  return 'Hello world!'

@app.route('/')
def index():
  return hello_world()
  
if __name__ == '__main__':
  app.run(host="0.0.0.0", port=8080)
