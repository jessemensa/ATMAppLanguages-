
from Transaction import Transaction


class BalanceInquiry(Transaction): 
    def __init__(self, useraccountnumber, atmscreen, atmbankdatabase):
        super().__init__(useraccountnumber, atmscreen, atmbankdatabase)

    def execute(self):
        bankdatabase = self.getbankdatabase()
        screen = self.getscreen() 
        availablebalance = bankdatabase.getavailablebalance(self.getaccountnumber())
        totalbalance = bankdatabase.gettotalbalance(self.getaccountnumber()) 

        # display the balance information on the screen 
        screen.display_message_line("\nBalance Information:")
        screen.display_message(" - Available balance: ")
        screen.display_dollar_amount(availablebalance)
        screen.display_message("\n - Total balance: ")
        screen.display_dollar_amount(totalbalance)
        screen.display_message_line("")