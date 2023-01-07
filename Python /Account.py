class Account: 
    def __init__(self, the_account_number, the_pin, the_available_balance, the_total_balance):
        self.account_number = the_account_number; 
        self.pin = the_pin; 
        self.available_balance = the_available_balance; 
        self.total_balance = the_total_balance;

    def validate_pin(self, user_pin):
        if user_pin == self.pin:
            return True; 
        else:
            return False;

    def get_available_balance(self):
        return self.available_balance; 
    
    def get_total_balance(self):
        return self.total_balance; 

    def credit(self, amount):
        self.total_balance += amount; 

    def debit(self, amount):
        self.available_balance -= amount; 
        self.total_balance -= amount;

    def get_account_number(self):
        return self.account_number;