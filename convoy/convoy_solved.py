import sqlite3

def main():
    con = sqlite3.connect(":memory:")
    cur = con.cursor()
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
    cur.executemany("INSERT into convoy(shipper_id, shipment_id, arrival, departure) VALUES(?, ?, ?, ?)", data)
    # get each shipper_id with '_CA' appended to the name if they have any shipments to or from CA
    # get total number of shipments per shipper alonside
    # order by total descending and then by shipper_id ascending
    query = """
        SELECT
            CASE WHEN SUM(CASE WHEN arrival = 'CA' OR departure = 'CA' THEN 1 ELSE 0 END) THEN shipper_id || '_CA' ELSE shipper_id END as shipper_id,
            COUNT(*) AS TOTAL
        FROM convoy
        GROUP BY shipper_id
        ORDER BY COUNT(*) DESC, shipper_id ASC
    """
    for row in cur.execute(query):
        print(row)
    cur.close()
    con.close() # do not commit this temp data

if __name__ == "__main__":
    main()
    


