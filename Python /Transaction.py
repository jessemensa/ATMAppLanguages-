from abc import ABC, abstractmethod

class Transaction(ABC): 
    def __init__(self, useraccountnumber, atmscreen, atmbankdatabase):
        self.useraccountnumber = useraccountnumber
        self.atmscreen = atmscreen
        self.atmbankdatabase = atmbankdatabase

    def getaccountnumber(self):
        return self.useraccountnumber 

    def getscreen(self):
        return self.atmscreen

    def getbankdatabase(self):
        return self.atmbankdatabase 

    @abstractmethod 
    def execute(self): 
        pass 