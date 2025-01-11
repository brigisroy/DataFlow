import csv
import random
from datetime import datetime, timedelta

# Helper function to generate random data
def generate_random_data(num_records=1000):
    categories = ['Shoes', 'Electronics', 'Clothing']
    regions = ['North America', 'Europe', 'Asia', 'South America']
    payment_methods = ['Credit Card', 'PayPal', 'Debit Card']
    product_names = ['UltraBoost Running Shoes', 'iPhone 15 Pro', 'Levi\'s 501 Jeans', 'Sony WH-1000XM5 Headphones']

    records = []
    for _ in range(num_records):
        order_id = random.randint(1000, 9999)
        product_id = f'P{random.randint(100, 999)}'
        customer_id = f'C{random.randint(100, 999)}'
        product_name = random.choice(product_names)
        category = random.choice(categories)
        region = random.choice(regions)
        date_of_sale = (datetime.now() - timedelta(days=random.randint(0, 365))).strftime('%Y-%m-%d')
        quantity_sold = random.randint(1, 5)
        unit_price = round(random.uniform(50, 1500), 2)
        discount = round(random.uniform(0, 0.2), 2)
        shipping_cost = round(random.uniform(5, 30), 2)
        payment_method = random.choice(payment_methods)
        customer_name = f'Customer_{random.randint(1, 100)}'
        customer_email = f'{customer_name.lower()}@email.com'
        customer_address = f'{random.randint(1, 999)} Main St, City, State {random.randint(10000, 99999)}'

        record = [
            order_id, product_id, customer_id, product_name, category, region, date_of_sale,
            quantity_sold, unit_price, discount, shipping_cost, payment_method, customer_name,
            customer_email, customer_address
        ]
        records.append(record)

    return records

# Generate records
records = generate_random_data(1000)

# Define CSV filename
filename = '/data/generated_orders.csv'

# Write to CSV
header = [
    "Order ID", "Product ID", "Customer ID", "Product Name", "Category", "Region",
    "Date of Sale", "Quantity Sold", "Unit Price", "Discount", "Shipping Cost",
    "Payment Method", "Customer Name", "Customer Email", "Customer Address"
]

with open(filename, mode='w', newline='', encoding='utf-8') as file:
    writer = csv.writer(file)
    writer.writerow(header)
    writer.writerows(records)

filename