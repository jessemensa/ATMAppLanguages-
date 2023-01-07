from Transaction import Transaction


class Deposit(Transaction):
    def __init__(self, user_account_number, atm_screen, atm_bank_database, atm_keypad, atm_deposit_slot):
        super().__init__(user_account_number, atm_screen, atm_bank_database)
        self.keypad = atm_keypad
        self.deposit_slot = atm_deposit_slot
        self.amount = 0
        self.CANCELED = 0

    def execute(self):
        bank_database = self.get_bank_database()
        screen = self.get_screen()

        self.amount = self.prompt_for_deposit_amount()

        if self.amount != self.CANCELED:
            screen.display_message("\nPlease insert a deposit envelope containing ")
            screen.display_dollar_amount(self.amount)
            screen.display_message_line(".")

            envelope_received = self.deposit_slot.is_envelope_received()

            if envelope_received:
                screen.display_message_line("\nYour envelope has been received. \nNOTE: The money just deposited will not be available until we verify the amount of any enclosed cash and your checks clear.")
                bank_database.credit(self.get_account_number(), self.amount)
            else:
                screen.display_message_line("\nYou did not insert an envelope, so the ATM has canceled your transaction.")
        else:
            screen.display_message_line("\nCanceling transaction...")

    def prompt_for_deposit_amount(self):
        screen = self.get_screen()

        screen.display_message("\nPlease enter a deposit amount in CENTS (or 0 to cancel): ")
        input_ = self.keypad.get_input()

        if input_ == self.CANCELED:
            return self.CANCELED
        else:
            return input_ / 100
