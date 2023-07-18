from flask import Flask, jsonify, request
import requests

orders = []

app = Flask(__name__)

@app.route('/orders', methods=["GET"])
def list_orders():
    return {"orders": orders}

@app.route('/orders', methods=["POST"])
def create_order():
    order_data = request.get_json()
    if not order_data:
        return jsonify({'message': 'Invalid JSON format'}), 400

    print(order_data)

    user_id = order_data.get('user_id')
    cart = order_data.get('cart')

    print(user_id)
    print(cart)

    if user_id is None or cart is None:
        return jsonify({'message': 'Missing required fields'}), 400

    user_response = requests.get(f'http://user-management/users/{user_id}')
    if user_response.status_code == 404:
        return jsonify({'message': 'User not found'}), 404

    order = {
        'user_id': user_id,
        'cart': cart
    }
    orders.append(order)
    return ""

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=80)
