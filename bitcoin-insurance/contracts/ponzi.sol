contract Government {

    // Global Variables
    uint32 public lastCreditorPayedOut;
    address[] public creditorAddresses;
    uint[] public creditorAmounts;
    address public corruptElite;
    mapping (address => uint) buddies;
    uint8 public round;

    function Government() {
        // The corrupt elite establishes a new government
        // this is the commitment of the corrupt Elite - everything that can not be saved from a crash
        corruptElite = msg.sender;
    }

    function lendGovernmentMoney(address buddy) returns (bool) {
        uint amount = msg.value;
        if (amount > 0) {
            // register the new creditor and his amount with 10% interest rate
            creditorAddresses.push(msg.sender);
            creditorAmounts.push(amount * 110 / 100);
            // now the money is distributed
            // firstly the corrupt elite grabs 1% - thieves!
            corruptElite.send(amount * 1/100);
            // if you have a buddy in the government (and he is in the creditor list) he can get 5% of your credits.
            // Make a deal with him.
            if(buddies[buddy] >= amount) {
                buddy.send(amount * 2/100);
            }
            buddies[msg.sender] += amount * 110 / 100;
            // 97.5% of the money will be used to pay out old creditors
            // if (creditorAmounts[lastCreditorPayedOut] <= address(this).balance - profitFromCrash) {
                // creditorAddresses[lastCreditorPayedOut].send(creditorAmounts[lastCreditorPayedOut]);
                // buddies[creditorAddresses[lastCreditorPayedOut]] -= creditorAmounts[lastCreditorPayedOut];
                lastCreditorPayedOut += 1;
            // }
            return true;
        }
        else {
            msg.sender.send(amount);
            return false;
        }
    }

    // fallback function
    function() {
        lendGovernmentMoney(0);
    }

    function totalDebt() returns (uint debt) {
        for(uint i=lastCreditorPayedOut; i<creditorAmounts.length; i++){
            debt += creditorAmounts[i];
        }
    }

    function totalPayedOut() returns (uint payout) {
        for(uint i=0; i<lastCreditorPayedOut; i++){
            payout += creditorAmounts[i];
        }
    }

    // From time to time the corrupt elite inherits it's power to the next generation
    function usurpation() {
        if (msg.value >= 1 * 10 ** 18) {
            corruptElite = msg.sender;
        }
    }

    function getCreditorAddresses() returns (address[]) {
        return creditorAddresses;
    }

    function getCreditorAmounts() returns (uint[]) {
        return creditorAmounts;
    }
}
