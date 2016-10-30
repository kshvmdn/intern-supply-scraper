import csv
import json


def _json(fn, data):
    with open(fn, 'w') as f:
        f.write(json.dumps(data))

    return True


def _csv(fn, data):
    with open(fn, 'w') as f:
        writer = csv.DictWriter(f, data[0].keys(), lineterminator='\n')

        writer.writeheader()

        for d in data:
            writer.writerow(d)

    return True
