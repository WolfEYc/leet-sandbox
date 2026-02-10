import sqlite3

def main():
    con = sqlite3.connect(":memory:")
    with con.cursor() as cur:
        cur.execute("CREATE TABLE convoy(shipper_id TEXT, shipment_id INTEGER, arrival TEXT, departure TEXT)")
        data = [
            ("walmart", 1914, "TX", "AZ"),
            ("amazon", 33892, "CA", "NY"),
            ("amazon", 98292, "TX", "NM"),
            ("walmart", 3982, "TN", "VA"),
            ("HEB", 999298, "TX", "TX"),
            ("whole-foods", 982282, "MN", "IL"),
            ("HEB", 22922, "ND", "SD"),
            ("whole-foods", 18228, "AK", "CA"),
            ("trader-joes", 928383, "MN", "NV"),
            ("amazon", 8938383, "OH", "UT"),
        ]
        cur.execute("INSERT into convoy(shipper_id, shipment_id, arrival, departure) VALUES(?)", data)
        # get each shipper_id with '_CA' appended to the name if they have any shipments to or from CA
        # get total number of shipments per shipper alonside
        query = """
        
        """
        for row in cur.execute(query):
            print(row)
    con.close() # do not commit this temp data
    


