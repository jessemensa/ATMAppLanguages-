

from Transaction import Transaction


class Withdrawal(Transaction):
    def __init__(self, useraccountnumber, atmscreen, atmbankdatabase, atmkeypad, atmcashdispenser):
        super().__init__(useraccountnumber, atmscreen, atmbankdatabase)
        self.atmkeypad = atmkeypad
        self.atmcashdispenser = atmcashdispenser
        self.amount = 0 
        self.CANCELED = 6; 

    def execute(self):
        cashdispensed = False 
        while not cashdispensed: 
            self.amount = self.display_menu_of_amounts()
            if self.amount != self.CANCELED:
                available_balance = self.get_bank_database().get_available_balance(self.get_account_number())

                if self.amount <= available_balance:
                    if self.cash_dispenser.is_sufficient_cash_available(self.amount):
                        self.get_bank_database().debit(self.get_account_number(), self.amount)

                        self.cash_dispenser.dispense_cash(self.amount)
                        cash_dispensed = True

                        self.get_screen().display_message_line("\nYour cash has been dispensed. Please take your cash now.")
                    else:
                        self.get_screen().display_message_line(
                            "\nInsufficient cash available in the ATM. \nPlease choose a smaller amount.")
                else:
                    self.get_screen().display_message_line("\nInsufficient funds in your account. \nPlease choose a smaller amount.")
            else:
                self.get_screen().display_message_line("\nCanceling transaction...")
                return
    
    def display_menu_of_amounts(self):
        userchoice = 0 
        amounts = [0, 20, 40, 60, 100, 200] 
        while userchoice == 0: 
            self.getscreen.display_message_line("\nWithdrawal Menu:") 
            self.getscreen.display_message_line("1 - £20")
            self.getscreen.display_message_line("2 - £40")
            self.getscreen.display_message_line("3 - £60")
            self.getscreen.display_message_line("4 - £100")
            self.getscreen.display_message_line("5 - £200")
            self.getscreen.display_message_line("6 - Cancel transaction")

            input = self.getkeypad.getinput() 

            if input in range(1, 6):
                userchoice = amounts[input]
            elif input == self.CANCELED:
                userchoice = self.CANCELED
            else: 
                self.getscreen.display_message_line("\nInvalid selection. Try again.")
            return userchoice 