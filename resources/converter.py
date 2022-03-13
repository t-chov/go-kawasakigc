import csv
import sys

if __name__ == '__main__':
    garbages = set()
    with open(sys.argv[1]) as csvfile:
        reader = csv.DictReader(csvfile)
        writer = csv.DictWriter(sys.stdout, [
            'name',
            'GarbageType',
            'DetailType',
            'Description',
            'Url',
        ])
        writer.writeheader()
        for row in reader:
            names = [row['品目名']]
            if row['類似語']:
                synonyms = row['類似語'].split(' ')
                names.extend(synonyms)
            for name in names:
                if name in garbages:
                    continue
                writer.writerow({
                    'name': name,
                    'GarbageType': row['出し方（一覧）'],
                    'DetailType': row['出し方（詳細）'],
                    'Description': row['出し方のポイント'],
                    'Url': row['URL1'],
                })
                garbages.add(name)
