from flask import Flask

app = Flask(__name__)

@app.route('/healthz')
def healthz():
    return 'ok'

@app.route('/')
def index():
    return 'Hello from Python!'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
