import Account


class BankDatabase: 
    def __init__(self):
        self.accounts = [Account(12345, 54321, 1000.0, 1200.0),
                         Account(98765, 56789, 200.0, 200.0)]

    def get_account(self, account_number):
        for current_account in self.accounts:
            if current_account.get_account_number() == account_number:
                return current_account; 
        return None;

    def authenticate_user(self, account_number, pin):
        account = self.get_account(account_number)

        if account is not None:
            return account.validate_pin(pin)
        else:
            return False

    def get_available_balance(self, account_number):
        return self.get_account(account_number).get_available_balance() 

    def get_total_balance(self, account_number):
        return self.get_account(account_number).get_total_balance() 

    def credit(self, account_number, amount):
        self.get_account(account_number).credit(amount) 

    def debit(self, account_number, amount):
        self.get_account(account_number).debit(amount)