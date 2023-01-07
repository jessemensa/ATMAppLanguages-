
class CashDispenser: 
    INITIAL_COUNT = 500; 
    def __init__(self):
        self.count = self.INITIAL_COUNT; 

    def dispense_cash(self, amount):
        bills_required = amount // 20; 
        self.count -= bills_required; 

    def is_sufficient_cash_available(self, amount):
        bills_required = amount // 20; 
        if self.count >= bills_required:
            return True; 
        else:
            return False;