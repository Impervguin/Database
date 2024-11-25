import subprocess
import os
import time
import hashlib
from datetime import datetime, date, timedelta
import faker
import random as r

EMAILS = set()
PHONES = set()

class Client:
    def __init__(self, firstName="", lastName="", dateOfBirth=date.today(), email="", phone="", address="", createdAt=datetime.now()) -> None:
        self.firstName = firstName
        self.lastName = lastName
        self.dateOfBirth = dateOfBirth
        self.email = email
        self.phone = phone
        self.address = address
        self.createdAt = createdAt
    
    def RandomFillSelf(self, startDate, endDate):
        fack = faker.Faker()
        self.firstName = fack.first_name()
        self.lastName = fack.last_name()
        self.dateOfBirth = fack.date_of_birth(minimum_age=21, maximum_age=70)
        self.email = fack.email()
        while self.email in EMAILS:
            self.email = fack.email()
        EMAILS.add(self.email)
        self.phone = r.choice(["+", ""]) + str(r.randint(1, 9)) + fack.basic_phone_number().replace('(', '').replace(')', '')
        while self.phone in PHONES:
            self.phone = r.choice(["+", ""]) + str(r.randint(1, 9)) + fack.basic_phone_number().replace('(', '').replace(')', '')
        PHONES.add(self.phone)
        self.address = fack.address()
        self.createdAt = fack.date_between_dates(date_start=startDate, date_end=endDate)
    
    @staticmethod
    def GetCsvHeader():
        return "first_name,last_name,dob,email,phone_number,address,created_at"

    def GetCsvString(self):
        return ",".join(map(str, (self.firstName, self.lastName, self.dateOfBirth, self.email, '"'+self.phone+'"', "'"+self.address.replace("\n", " ").replace(",", ";")+"'", self.createdAt)))


CLIENT_PER_FILE = 100
FILE_PER_SECONDS = 300
FILE_DIRECTORY = 'deployments/nifi/input_file'

if __name__ == "__main__":
    startDate = datetime(2000, 1, 1)
    endDate = datetime.now()
    while True:
        pid = os.fork()
        if pid:
            time.sleep(FILE_PER_SECONDS)
        else:
            pid = os.getpid()
            print(f"Process {pid} started")
            fname = f'client_{hashlib.md5(str(pid).encode("utf8")).hexdigest()}_{datetime.today().strftime('%Y-%m-%d_%H:%M:%S')}.csv'
            with open(os.path.join(FILE_DIRECTORY, fname), 'w') as file:
                file.write(Client.GetCsvHeader() + "\n")
                for i in range(CLIENT_PER_FILE):
                    client = Client()
                    client.RandomFillSelf(startDate, endDate)
                    file.write(client.GetCsvString() + "\n")
            print(f"Process {pid} finished")
            os._exit(0)
    
