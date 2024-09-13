from datetime import datetime, date, time, timedelta
import datetime as dt
import dateutil as du
import faker
import random as r
import hashlib
import os

EMAILS = set()
PHONES = set()
USERNAMES = set()

class Client:
    def __init__(self, id=0, firstName="", lastName="", dateOfBirth=date.today(), email="", phone="", address="", createdAt=datetime.now()) -> None:
        self.id = id
        self.firstName = firstName
        self.lastName = lastName
        self.dateOfBirth = dateOfBirth
        self.email = email
        self.phone = phone
        self.address = address
        self.createdAt = createdAt
    
    def RandomFillSelf(self, startDate, endDate, id=None):
        if id is not None:
            self.id = id
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
        return "id,first_name,last_name,dob,email,phone_number,address,created_at"

    def GetCsvString(self):
        return ",".join(map(str, (self.id, self.firstName, self.lastName, self.dateOfBirth, self.email, self.phone, "'"+self.address.replace("\n", " ").replace(",", ";")+"'", self.createdAt)))

class Account:
    TYPES = {"savings":(2, 3), "checking": (4, 8), "fd": (9, 17), "credit": (16, 22)}
    STATUSES = ("active", "inactive", "closed")
    def __init__(self, id=0, clientID=0, type="", balance=.0, interest=.0, status="", createdAt=date.today()) -> None:
        self.id = id
        self.clientID = clientID
        self.type = type
        self.balance = balance
        self.interest = interest
        self.status = status
        self.createdAt = createdAt
    
    def RandomFillSelf(self, client: Client, endDate, id=None):
        if id is not None:
            self.id = id
        self.clientID = client.id
        fack = faker.Faker()
        self.type = r.choice(list(self.TYPES.keys()))
        self.balance = r.randint(1000, 1000000)
        self.interest = self.TYPES[self.type][0] + (r.random() + 0.05) * (self.TYPES[self.type][1] - self.TYPES[self.type][0])
        self.status = r.choice(self.STATUSES)
        self.createdAt = fack.date_between_dates(client.createdAt, endDate)

    @staticmethod
    def GetCsvHeader():
        return "id,client_id,balance,interest,created_at,atype,astatus"
    
    def GetCsvString(self):
        return ",".join(map(str, (self.id, self.clientID, self.balance, self.interest, self.createdAt, self.type, self.status)))
        
class Card:
    STATUSES = ("active", "blocked", "expired")
    def __init__(self, id=0, accountId=0, number="", cvv="", expiredAt=date.today(), createdAt=date.today(), status=""):
        self.id = id
        self.accountId = accountId
        self.number = number
        self.cvv = cvv
        self.expiredAt = expiredAt
        self.createdAt = createdAt
        self.status = status
    
    def RandomFillSelf(self, account: Account, endDate, id=None):
        if id is not None:
            self.id = id
        self.accountId = account.id
        fack = faker.Faker()
        self.status = r.choice(self.STATUSES)
        self.createdAt = fack.date_between_dates(account.createdAt, endDate - timedelta(days=1))
        if self.status == "expired":
            self.expiredAt = fack.date_between_dates(self.createdAt + timedelta(days=1), endDate)
        else:
            self.expiredAt = fack.date_between_dates(endDate, endDate + timedelta(days=365*5))
        self.number = "".join(map(str, [r.randint(1, 9) for _ in range(16)]))
        self.cvv = str(r.randint(1, 9)) + str(r.randint(1, 9)) + str(r.randint(1, 9))
    
    @staticmethod
    def GetCsvHeader():
        return "id,account_id,cnumber,cvv,created_at,expired_at,cstatus"

    def GetCsvString(self):
        return ",".join(map(str, (self.id, self.accountId, self.number, self.cvv, self.createdAt, self.expiredAt, self.status)))

class Loan:
    STATUSES = ("active", "closed", "defaulted")
    def __init__(self, id=0, customerId=None, amount=.0, interest=.0, monthlyPayment=0., startDate=date.today(), endDate=date.today(), status="", description=""):
        self.id = id
        self.customerId = customerId
        self.amount = amount
        self.remaining = 0
        self.interest = interest
        self.monthlyPayment = monthlyPayment
        self.startDate = startDate
        self.endDate = endDate
        self.status = status
        self.description = description

    def RandomFillSelf(self, client: Client, endDate=date.today(), id=None):
        if id is not None:
            self.id = id
        self.customerId = client.id
        fack = faker.Faker()
        self.status = r.choice(self.STATUSES)
        # self.status = "active"
        self.amount = r.randint(0, 10000000)
        self.interest = r.randint(1, 45)
        self.description = ""
        self.startDate = fack.date_between_dates(client.createdAt, endDate)
        if self.status == "active":
            self.endDate = fack.date_between_dates(endDate, endDate + timedelta(days=365*20))
        else:
            self.remaining = 0
            self.endDate = fack.date_between_dates(self.startDate, endDate) + timedelta(days=31)
        self.endDate = self.endDate.replace(day=1, year=self.endDate.year, month=self.endDate.month)
        monthsDiff = self.endDate.month - self.startDate.month
        yearsDiff = self.endDate.year - self.startDate.year
        paymonths = yearsDiff * 12 + monthsDiff
        self.monthlyPayment = self.amount * (self.interest / 100) / 12 * (1 + (self.interest / 100) / 12) ** paymonths / ((1 + (self.interest / 100) / 12) ** paymonths - 1)
        if self.status == "active":
            monthsPassed = endDate.month - self.startDate.month + 12 * (endDate.year - self.startDate.year) 
            self.remaining = (paymonths - monthsPassed) * self.monthlyPayment
    
    @staticmethod
    def GetCsvHeader():
        return "id,client_id,amount,interest,remaining_amount,monthly_payment,start_date,end_date,lstatus,ldescription"
    
    def GetCsvString(self):
        return ",".join(map(str, (self.id, self.customerId, self.amount, self.interest, self.remaining, self.monthlyPayment, self.startDate, self.endDate, self.status, self.description)))
        
class Transaction:
    TYPES = ("deposit", "withdrawal", "transfer")
    def __init__(self, id, accountId, amount, type, doneAt, balanceAfter, systemDescription, userDescription):
        self.id = id
        self.accountId = accountId
        self.amount = amount
        self.type = type
        self.doneAt = doneAt
        self.balanceAfter = balanceAfter
        self.systemDescription = systemDescription
        self.userDescription = userDescription
    
    @staticmethod
    def GetCsvHeader():
        return "id,account_id,ttype,amount,done_at,balance_after,system_description,client_description"
    
    def GetCsvString(self):
        return ",".join(map(str, (self.id, self.accountId, self.type, self.amount, self.doneAt, self.balanceAfter, self.systemDescription.replace("\n", ""), self.userDescription.replace("\n", ""))))

class Service:
    def __init__(self, id=0, name="", description="", fee=0) -> None:
        self.id = id
        self.name = name
        self.description = description
        self.fee = fee
    
    def GenerateRandomSelf(self, id=None):
        if id is not None:
            self.id = id
        fack = faker.Faker()
        self.name = fack.catch_phrase()
        self.description = fack.text(max_nb_chars=100).replace("\n", " ").replace(",", " ")
        self.fee = fack.random_int(1, 9999)
    
    @staticmethod
    def GetCsvHeader():
        return "id,sname,sdescription,fee"
    
    def GetCsvString(self):
        return ",".join(map(str, (self.id, self.name, "'"+self.description+"'", self.fee)))

class ClientService:
    def __init__(self, id=0, clientId=0, serviceId=0) -> None:
        self.id = id
        self.clientId = clientId
        self.serviceId = serviceId
    
    @staticmethod
    def GetCsvHeader():
        return "id,client_id,service_id"
    
    def GetCsvString(self):
        return ",".join(map(str, (self.id, self.clientId, self.serviceId)))

class User:
    STATUSES = ("active", "blocked")
    def __init__(self, id=0, username="", hashPassword="", lastLogin=date.today(), failesAttempts=0, status=''):
        self.id = id
        self.clientId = 0
        self.username = username
        self.hashPassword = hashPassword
        self.lastLogin = lastLogin
        self.failesAttempts = failesAttempts
        self.status = status
    
    def GenerateRandomSelf(self, client, id=0):
        if id is not None:
            self.id = id
        fack = faker.Faker()
        self.clientId = client.id
        self.username = fack.user_name()
        while self.username in USERNAMES:
            self.username = fack.user_name()
        USERNAMES.add(self.username)
        self.hashPassword = hashlib.sha256(fack.password(length=16).encode()).hexdigest()
        self.lastLogin = fack.past_date()
        self.failesAttempts = 0
        self.status = r.choice(self.STATUSES)

    @staticmethod
    def GetCsvHeader():
        return "id,client_id,username,hashpassword,last_login,failes_attempts,ustatus"
    
    def GetCsvString(self):
        return ",".join(map(str, (self.id,self.clientId, self.username, self.hashPassword, self.lastLogin, self.failesAttempts, self.status)))
    
class Notification:
    def __init__(self, clientId, message, sentAt, read):
        self.clientId = clientId
        self.message = message
        self.sentAt = sentAt
        self.read = read
    
    @staticmethod
    def GetCsvHeader():
        return "client_id,nmessage,created_at,seen"
    
    def GetCsvString(self):
        return ",".join(map(str, (self.clientId, self.message, self.sentAt, self.read)))
        

class iBankSimulation:
    def __init__(
            self,
            clientsCount=1000,
            accountsCount=1000,
            cardsCount=2000,
            loansCount=100,
            transactionCount=10000,
            servicesCount=100,
            clientServicesCount=1000,
            ):
        self.clientsCount=clientsCount
        self.accountsCount=accountsCount
        self.cardsCount=cardsCount
        self.loansCount=loansCount
        self.transactionCount=transactionCount
        self.servicesCount=servicesCount
        self.clientServicesCount=clientServicesCount
        self.clients = []
        self.accounts = []
        self.cards = []
        self.loans = []
        self.transactions = []
        self.services = []
        self.clientServices = []
        self.users = []
        self.notifications = []
    
    def generateClients(self, startDate, endDate):
        for i in range(self.clientsCount):
            client = Client()
            client.RandomFillSelf(startDate, endDate, id=i)
            self.clients.append(client)
    
    def generateAccounts(self, endDate):
        for i in range(self.accountsCount):
            account = Account()
            account.RandomFillSelf(self.clients[r.randint(0, self.clientsCount-1)], endDate,id=i)
            self.accounts.append(account) 
    
    def generateCards(self, endDate):
        for i in range(self.cardsCount):
            card = Card()
            card.RandomFillSelf(self.accounts[r.randint(0, self.accountsCount-1)], endDate, id=i)
            self.cards.append(card)
    
    def generateLoans(self, endDate):
        for i in range(self.loansCount):
            loan = Loan()
            loan.RandomFillSelf(self.clients[r.randint(0, self.clientsCount-1)], endDate, id=i)
            self.loans.append(loan)
    
    def generateServices(self):
        for i in range(self.servicesCount):
            service = Service()
            service.GenerateRandomSelf(id=i)
            self.services.append(service)
    
    def generateServiceClient(self):
        for i in range(self.clientServicesCount):
            clientService = ClientService(id=i, clientId = r.randint(0, self.clientsCount-1), serviceId = r.randint(0, self.servicesCount-1))
            self.clientServices.append(clientService)
    
    def generateUsers(self):
        for i in range(self.clientsCount):
            user = User()
            user.GenerateRandomSelf(self.clients[i], id=i)
            self.users.append(user)
    
    def Simulate(self, startDate=date.today() - timedelta(days=365), endDate=date.today(), verbose=False):
        self.generateClients(startDate, endDate)
        if verbose: print("Clients created")
        self.generateAccounts(endDate)
        if verbose: print("Accounts created")
        self.generateCards(endDate)
        if verbose: print("Cards created")
        self.generateLoans(endDate)
        if verbose: print("Loans created")
        self.generateServices()
        if verbose: print("Services created")
        self.generateServiceClient()
        if verbose: print("Client Services created")
        self.generateUsers()
        if verbose: print("Users created")

        days = endDate - startDate
        transactionPerDay = self.transactionCount / days.days
        transactionCount = 0
        transactionId = 0
        fack = faker.Faker()
        for day in range(days.days + 1):
            currentDay = startDate + timedelta(days=day)
            if verbose:
                print(currentDay)
            transactionCount += transactionPerDay
            if currentDay.day == 1:
                for loanId in range(self.loansCount):
                    loan = self.loans[loanId]
                    if loan.startDate <= currentDay <= loan.endDate:
                        for account in self.accounts:
                            if account.createdAt <= currentDay and account.clientID == loan.customerId:
                                cAccount = account
                                break
                        else:
                            continue
                        if cAccount.balance < loan.monthlyPayment:
                            plus = r.randint(int(loan.monthlyPayment), 2 * int(loan.monthlyPayment))
                            transaction = Transaction(
                                id=transactionId,
                                accountId=cAccount.id,
                                amount=plus,
                                type="deposit",
                                doneAt=currentDay,
                                balanceAfter=cAccount.balance + plus,
                                systemDescription="",
                                userDescription="For loan payment"
                            )
                            self.transactions.append(transaction)
                            transactionCount -= 1
                            cAccount.balance += plus
                            transactionId += 1
                        minus = loan.monthlyPayment
                        transaction = Transaction(
                            id=transactionId,
                            accountId=cAccount.id,
                            amount=-loan.monthlyPayment,
                            type="transfer",
                            doneAt=currentDay,
                            balanceAfter=cAccount.balance - minus,
                            systemDescription="",
                            userDescription="Loan payment"
                        )
                        self.transactions.append(transaction)
                        transactionCount -= 1
                        cAccount.balance -= minus
                        transactionId += 1
            while transactionCount > 0:
                types = ["transfer", "deposit", "withdraw"]
                type = r.choice(types)
                if type == "transfer":
                    accountId1, accountId2 = r.randint(0, self.accountsCount - 1), r.randint(0, self.accountsCount - 1)
                    account1, account2 = self.accounts[accountId1], self.accounts[accountId2]
                    if account1.clientID == account2.clientID:
                        continue
                    amount = r.randint(0, int(account1.balance))
                    transaction = Transaction(
                            id=transactionId,
                            accountId=account1.id,
                            amount=-amount,
                            type="transfer",
                            doneAt=currentDay,
                            balanceAfter=account1.balance - amount,
                            systemDescription="",
                            userDescription=fack.text()
                        )
                    self.transactions.append(transaction)
                    transactionCount -= 1
                    transactionId += 1
                    account1.balance -= amount
                    transaction = Transaction(
                            id=transactionId,
                            accountId=account2.id,
                            amount=amount,
                            type="transfer",
                            doneAt=currentDay,
                            balanceAfter=account1.balance + amount,
                            systemDescription="",
                            userDescription=fack.text()
                        )
                    self.transactions.append(transaction)
                    transactionCount -= 1
                    transactionId += 1
                    account1.balance += amount
                elif type == "deposit":
                    accountId = r.randint(0, self.accountsCount - 1)
                    account = self.accounts[accountId]
                    amount = r.randint(0, int(account.balance))
                    transaction = Transaction(
                            id=transactionId,
                            accountId=account.id,
                            amount=amount,
                            type="deposit",
                            doneAt=currentDay,
                            balanceAfter=account.balance + amount,
                            systemDescription="",
                            userDescription=fack.text()
                        )
                    self.transactions.append(transaction)
                    transactionCount -= 1
                    transactionId += 1
                    account.balance += amount
                else:
                    accountId = r.randint(0, self.accountsCount - 1)
                    account = self.accounts[accountId]
                    amount = r.randint(0, int(account.balance))
                    if account.balance < amount:
                        continue
                    transaction = Transaction(
                            id=transactionId,
                            accountId=account.id,
                            amount=-amount,
                            type="withdraw",
                            doneAt=currentDay,
                            balanceAfter=account.balance - amount,
                            systemDescription="",
                            userDescription=fack.text()
                        )
                    self.transactions.append(transaction)
                    transactionCount -= 1
                    transactionId += 1
                    account.balance -= amount
            





            
    def exportcsv(self, folder="./ibank"):
        with open(os.path.join(folder, "client.csv"), mode="w") as csvfile:
            csvfile.write(Client.GetCsvHeader() + "\n")
            for client in self.clients:
                csvfile.write(client.GetCsvString() + "\n")
        
        with open(os.path.join(folder, "account.csv"), mode="w") as csvfile:
            csvfile.write(Account.GetCsvHeader() + "\n")
            for account in self.accounts:
                csvfile.write(account.GetCsvString() + "\n")
        
        with open(os.path.join(folder, "card.csv"), mode="w") as csvfile:
            csvfile.write(Card.GetCsvHeader() + "\n")
            for card in self.cards:
                csvfile.write(card.GetCsvString() + "\n")
        
        with open(os.path.join(folder, "loan.csv"), mode="w") as csvfile:
            csvfile.write(Loan.GetCsvHeader() + "\n")
            for loan in self.loans:
                csvfile.write(loan.GetCsvString() + "\n")
        
        with open(os.path.join(folder, "service.csv"), mode="w") as csvfile:
            csvfile.write(Service.GetCsvHeader() + "\n")
            for service in self.services:
                csvfile.write(service.GetCsvString() + "\n")
        
        with open(os.path.join(folder, "transaction.csv"), mode="w") as csvfile:
            csvfile.write(Transaction.GetCsvHeader() + "\n")
            for transaction in self.transactions:
                csvfile.write(transaction.GetCsvString() + "\n")
        
        with open(os.path.join(folder, "clientservice.csv"), mode="w") as csvfile:
            csvfile.write(ClientService.GetCsvHeader() + "\n")
            for clientService in self.clientServices:
                csvfile.write(clientService.GetCsvString() + "\n")
        
        with open(os.path.join(folder, "user.csv"), mode="w") as csvfile:
            csvfile.write(User.GetCsvHeader() + "\n")
            for user in self.users:
                csvfile.write(user.GetCsvString() + "\n")
        
        with open(os.path.join(folder, "notification.csv"), mode="w") as csvfile:
            csvfile.write(Notification.GetCsvHeader() + "\n")
            for notification in self.notifications:
                csvfile.write(notification.GetCsvString() + "\n")
        
        print("CSV files exported successfully")



if __name__ == "__main__":
    bank = iBankSimulation()
    bank.Simulate(startDate=date(2022, 1, 1), endDate=date(2023, 1, 1), verbose=True)
    bank.exportcsv()
