from flask import Flask

app = Flask(__name__)

@app.route('/orders')
def hello():
    return {"message": 'Hello, Orders!'}

if __name__ == '__main__':
    app.run(port=80)
