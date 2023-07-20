import json
from flask import Flask, jsonify, request
import requests

orders = []

app = Flask(__name__)

@app.route('/orders', methods=["GET"])
def list_orders():
    result = []
    print(orders)

    for order in orders:
        user = requests.get(f'http://user-management/users/{order["user_id"]}')
        cart = []

        for cart_item in order["cart"]:
            product = requests.get(f'http://product-catalog/products/{cart_item["product_id"]}')
            cart.append({
                "product": product.json(),
                "amount": cart_item["amount"]
            })

        result.append({
            "user": user.json(),
            "cart": cart,
        })

    return {"orders": result}

@app.route('/orders', methods=["POST"])
def create_order():
    order_data = request.get_json()
    if not order_data:
        return jsonify({'message': 'Invalid JSON format'}), 400

    user_id = order_data.get('user_id')
    cart = order_data.get('cart')

    if user_id is None or cart is None:
        return jsonify({'message': 'Missing required fields'}), 400

    user_response = requests.get(f'http://user-management/users/{user_id}')
    if user_response.status_code == 404:
        return jsonify({'message': 'User not found'}), 404

    stored_cart = []

    for product in cart:
        product_id = product.get('product_id')
        amount = product.get('amount')

        if product_id is None or amount is None:
            return jsonify({'message': 'Missing required fields'}), 400

        product_response = requests.get(f'http://user-management/users/{product_id}')
        if product_response.status_code == 404:
            return jsonify({'message': 'Product not found'}), 404

        stored_cart.append({
            "product_id": product_id,
            "amount": amount,
        })

    order = {
        'user_id': user_id,
        'cart': stored_cart
    }
    orders.append(order)
    return ""

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=80)
