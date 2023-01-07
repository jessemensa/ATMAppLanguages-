class Screen: 
    def display_message(self, message):
        print(message) 

    def display_message_line(self, message):
        print(message)

    def display_dollar_amount(self, amount):
        print(f"£{amount:, .2f}")