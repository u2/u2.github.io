// Factory "morphs" into a Pudding class.
// The reasoning is that calling load in each context
// is cumbersome.

(function() {

  var contract_data = {
    abi: [{"constant":true,"inputs":[],"name":"godBank","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"round","outputs":[{"name":"","type":"uint32"}],"type":"function"},{"constant":false,"inputs":[],"name":"protectKingdom","outputs":[{"name":"","type":"bool"}],"type":"function"},{"constant":false,"inputs":[],"name":"abdicate","outputs":[],"type":"function"},{"constant":true,"inputs":[],"name":"kingCost","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"madKing","outputs":[{"name":"","type":"address"}],"type":"function"},{"constant":true,"inputs":[],"name":"jesterBank","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"citizensAddresses","outputs":[{"name":"","type":"address"}],"type":"function"},{"constant":true,"inputs":[],"name":"kingBank","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"piggyBank","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"lastCollection","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"totalCitizens","outputs":[{"name":"","type":"uint32"}],"type":"function"},{"constant":false,"inputs":[],"name":"usurpation","outputs":[],"type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"citizensAmounts","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[],"name":"collectFee","outputs":[],"type":"function"},{"constant":true,"inputs":[],"name":"lastCitizenPaid","outputs":[{"name":"","type":"uint32"}],"type":"function"},{"constant":true,"inputs":[],"name":"amountInvested","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"jester","outputs":[{"name":"","type":"address"}],"type":"function"},{"constant":true,"inputs":[],"name":"amountAlreadyPaidBack","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"trueGods","outputs":[{"name":"","type":"address"}],"type":"function"},{"constant":true,"inputs":[],"name":"onThrone","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[],"name":"murder","outputs":[],"type":"function"},{"inputs":[],"type":"constructor"}],
    binary: "60606040526000805433600160a060020a031991821681178355600b8054600180549094168317909355426002819055600355670de0b6b3a7640000600455600c849055600d849055604060020a60e060020a0319909216680100000000000000009091021763ffffffff19169055610d8190819061007d90396000f3606060405236156101065760e060020a600035046311613fc98114610111578063146ca5311461011a5780631fdf6e0c14610130578063314e99a21461016b5780635841b9bf14610244578063816f3f4d1461024d5780639935b96814610266578063a2e7241c1461026f578063b117a3e8146102b5578063b7d5d4c0146102be578063bd6bbc31146102c7578063bd7b09e4146102d0578063ccf1ab9b146102df578063d466a03f1461030d578063d4d5d32a14610338578063d94395e21461037a578063d954cbcb14610391578063e2202a4d1461039a578063e684aa5c146103ac578063f113fccd146103b5578063f64c08b1146103c7578063fbeaebc6146103d0575b61046c610477610134565b61047b60065481565b61048d600b5460e060020a900463ffffffff1681565b61047b5b600034662386f26fc100008110156104c457604051600160a060020a033316908390839082818181858883f1506109139350505050565b61046c600b54604060020a9004600160a060020a0390811633919091161480156101a55750600054600160a060020a039081163390911614155b1561047957600b54600854604051604060020a909204600160a060020a0316916000919082818181858883f19350505050506064600460005054602802046005600050541115610c0457600454600b54604051604060020a909104600160a060020a03169160009160646028909202919091049082818181858883f1505060045460058054606460289390930292909204909103905550610c32915050565b61047b60045481565b6104a7600b54604060020a9004600160a060020a031681565b61047b60075481565b6104a760043560098054829081101561000257506000527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0154600160a060020a031681565b61047b60085481565b61047b60055481565b61047b60025481565b61048d600b5463ffffffff1681565b61046c600b543490604060020a9004600160a060020a0390811633919091161415610c7f5761046e8161094e565b61047b600435600a805482908110156100025750600052600080516020610d61833981519152015481565b61046c600054600160a060020a03908116339190911614156104795760405160008054600654600160a060020a03919091169282818181858883f15050505050565b61048d600b54640100000000900463ffffffff1681565b61047b600d5481565b6104a7600154600160a060020a031681565b61047b600c5481565b6104a7600054600160a060020a031681565b61047b60035481565b61046c67016345785d8a0000348190108015906103fd5750600154600160a060020a039081163390911614155b15610c7a57600754600154604051600160a060020a03919091169160009182818181858883f1505050600781905560018054600160a060020a03191633908117909155604051600160a060020a039091169250348490039082818181858883f19350505050506104778161094e565b005b60048054820190555b505b565b60408051918252519081900360200190f35b6040805163ffffffff929092168252519081900360200190f35b60408051600160a060020a03929092168252519081900360200190f35b68056bc75e2d6310000081111561050e57604051600160a060020a0333169060009068056bc75e2d630fffff1934019082818181858883f15068056bc75e2d631000009450505050505b426201518060026000505401101561059f57600b5463ffffffff16600114156105cf57600980546000198101908110156100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ae90920154600160a060020a031692916064605f909202919091049082818181858883f193505050505061080e565b600980546001810180835582818380158290116109fd578183600052602060002091820191016109fd9190610868565b600b5463ffffffff16600214156106ba57600980546000198101908110156100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ae9290920154600160a060020a031692916064603c909202919091049082818181858883f15050600980549093506001198101925082101590506100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ad90920154600160a060020a0316929160646023909202919091049082818181858883f193505050505061080e565b600b54600363ffffffff919091161061080e57600980546000198101908110156100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ae9290920154600160a060020a0316929160646032909202919091049082818181858883f15050600980549093506001198101925082101590506100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ad90920154600160a060020a031692916064601e909202919091049082818181858883f15050600980549093506002198101925082101590506100025750805460009182527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ac0154600554604051600160a060020a039290921692916064600f909202919091049082818181858883f150505050505b600580546006805460649284029290920490910190556000905560018054600160a060020a0319163317815560098054918201808255909190828183801582901161087c5781836000526020600020918201910161087c91905b808211156109135760008155600101610868565b5050506000928352506020909120018054600160a060020a03191633179055600a8054600181018083558281838015829011610917578183600052602060002091820191016109179190610868565b600580546064605a840204019055600b805460e060020a80820463ffffffff16600101027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9091161790555b5090565b50505060009283525060209091206064606e840204910155600b8054600163ffffffff82160163ffffffff199091161790556109c0815b42600255600d80548201905560078054606460058481029190910491820190925560088054820190558154019055610d2260085467016345785d8a0000901061047957604051600854600b54604060020a9004600160a060020a0316916000919082818181858883f150505060085550565b6108cb600654670de0b6b3a764000090106104795760405160008054600654600160a060020a03919091169282818181858883f150505060065550565b5050506000928352506020909120018054600160a060020a03191633179055600a8054600181018083558281838015829011610a4c57818360005260206000209182019101610a4c9190610868565b50505060009283525060209091206064606e840204910155600b8054600163ffffffff82160163ffffffff19909116179055610a878161094e565b5b60076000505460086000505460066000505460056000505430600160a060020a03163103030303600a600050600b60049054906101000a900463ffffffff1663ffffffff16815481101561000257600091909152600080516020610d618339815191520154108015610b0d5750600b5463ffffffff8181166401000000009092041611155b1561091357600b54600980549091640100000000900463ffffffff16908110156100025760009182527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af810154600a8054600160a060020a03929092169392909181101561000257908252604051600080516020610d61833981519152909101549082818181858883f15050600b54600a805490945064010000000090910463ffffffff169250821015905061000257600b8054600093909352600080516020610d61833981519152820154600c8054909101905567ffffffff000000001990921660019190910164010000000002179055610a88565b600b54604051600554604060020a909204600160a060020a03169160009182818181858883f1505050600555505b600b8054600054600160a060020a0316604060020a027bffffffffffffffffffffffffffffffffffffffff000000000000000019909116179055670de0b6b3a7640000600455565b610002565b42620d2f006003600050540111158015610ca457506064600460005054609602048110155b15610c7a57600b54600854604051604060020a909204600160a060020a0316916000919082818181858883f1505060068054606460058702040190555050506004819055600b80547bffffffffffffffffffffffffffffffffffffffff00000000000000001916604060020a3302179055426003556104778161094e565b61047760075467016345785d8a0000901061047957600754600154604051600160a060020a03919091169160009182818181858883f15050506007555056c65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8",
    unlinked_binary: "60606040526000805433600160a060020a031991821681178355600b8054600180549094168317909355426002819055600355670de0b6b3a7640000600455600c849055600d849055604060020a60e060020a0319909216680100000000000000009091021763ffffffff19169055610d8190819061007d90396000f3606060405236156101065760e060020a600035046311613fc98114610111578063146ca5311461011a5780631fdf6e0c14610130578063314e99a21461016b5780635841b9bf14610244578063816f3f4d1461024d5780639935b96814610266578063a2e7241c1461026f578063b117a3e8146102b5578063b7d5d4c0146102be578063bd6bbc31146102c7578063bd7b09e4146102d0578063ccf1ab9b146102df578063d466a03f1461030d578063d4d5d32a14610338578063d94395e21461037a578063d954cbcb14610391578063e2202a4d1461039a578063e684aa5c146103ac578063f113fccd146103b5578063f64c08b1146103c7578063fbeaebc6146103d0575b61046c610477610134565b61047b60065481565b61048d600b5460e060020a900463ffffffff1681565b61047b5b600034662386f26fc100008110156104c457604051600160a060020a033316908390839082818181858883f1506109139350505050565b61046c600b54604060020a9004600160a060020a0390811633919091161480156101a55750600054600160a060020a039081163390911614155b1561047957600b54600854604051604060020a909204600160a060020a0316916000919082818181858883f19350505050506064600460005054602802046005600050541115610c0457600454600b54604051604060020a909104600160a060020a03169160009160646028909202919091049082818181858883f1505060045460058054606460289390930292909204909103905550610c32915050565b61047b60045481565b6104a7600b54604060020a9004600160a060020a031681565b61047b60075481565b6104a760043560098054829081101561000257506000527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0154600160a060020a031681565b61047b60085481565b61047b60055481565b61047b60025481565b61048d600b5463ffffffff1681565b61046c600b543490604060020a9004600160a060020a0390811633919091161415610c7f5761046e8161094e565b61047b600435600a805482908110156100025750600052600080516020610d61833981519152015481565b61046c600054600160a060020a03908116339190911614156104795760405160008054600654600160a060020a03919091169282818181858883f15050505050565b61048d600b54640100000000900463ffffffff1681565b61047b600d5481565b6104a7600154600160a060020a031681565b61047b600c5481565b6104a7600054600160a060020a031681565b61047b60035481565b61046c67016345785d8a0000348190108015906103fd5750600154600160a060020a039081163390911614155b15610c7a57600754600154604051600160a060020a03919091169160009182818181858883f1505050600781905560018054600160a060020a03191633908117909155604051600160a060020a039091169250348490039082818181858883f19350505050506104778161094e565b005b60048054820190555b505b565b60408051918252519081900360200190f35b6040805163ffffffff929092168252519081900360200190f35b60408051600160a060020a03929092168252519081900360200190f35b68056bc75e2d6310000081111561050e57604051600160a060020a0333169060009068056bc75e2d630fffff1934019082818181858883f15068056bc75e2d631000009450505050505b426201518060026000505401101561059f57600b5463ffffffff16600114156105cf57600980546000198101908110156100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ae90920154600160a060020a031692916064605f909202919091049082818181858883f193505050505061080e565b600980546001810180835582818380158290116109fd578183600052602060002091820191016109fd9190610868565b600b5463ffffffff16600214156106ba57600980546000198101908110156100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ae9290920154600160a060020a031692916064603c909202919091049082818181858883f15050600980549093506001198101925082101590506100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ad90920154600160a060020a0316929160646023909202919091049082818181858883f193505050505061080e565b600b54600363ffffffff919091161061080e57600980546000198101908110156100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ae9290920154600160a060020a0316929160646032909202919091049082818181858883f15050600980549093506001198101925082101590506100025750805460009182526005546040517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ad90920154600160a060020a031692916064601e909202919091049082818181858883f15050600980549093506002198101925082101590506100025750805460009182527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7ac0154600554604051600160a060020a039290921692916064600f909202919091049082818181858883f150505050505b600580546006805460649284029290920490910190556000905560018054600160a060020a0319163317815560098054918201808255909190828183801582901161087c5781836000526020600020918201910161087c91905b808211156109135760008155600101610868565b5050506000928352506020909120018054600160a060020a03191633179055600a8054600181018083558281838015829011610917578183600052602060002091820191016109179190610868565b600580546064605a840204019055600b805460e060020a80820463ffffffff16600101027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9091161790555b5090565b50505060009283525060209091206064606e840204910155600b8054600163ffffffff82160163ffffffff199091161790556109c0815b42600255600d80548201905560078054606460058481029190910491820190925560088054820190558154019055610d2260085467016345785d8a0000901061047957604051600854600b54604060020a9004600160a060020a0316916000919082818181858883f150505060085550565b6108cb600654670de0b6b3a764000090106104795760405160008054600654600160a060020a03919091169282818181858883f150505060065550565b5050506000928352506020909120018054600160a060020a03191633179055600a8054600181018083558281838015829011610a4c57818360005260206000209182019101610a4c9190610868565b50505060009283525060209091206064606e840204910155600b8054600163ffffffff82160163ffffffff19909116179055610a878161094e565b5b60076000505460086000505460066000505460056000505430600160a060020a03163103030303600a600050600b60049054906101000a900463ffffffff1663ffffffff16815481101561000257600091909152600080516020610d618339815191520154108015610b0d5750600b5463ffffffff8181166401000000009092041611155b1561091357600b54600980549091640100000000900463ffffffff16908110156100025760009182527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af810154600a8054600160a060020a03929092169392909181101561000257908252604051600080516020610d61833981519152909101549082818181858883f15050600b54600a805490945064010000000090910463ffffffff169250821015905061000257600b8054600093909352600080516020610d61833981519152820154600c8054909101905567ffffffff000000001990921660019190910164010000000002179055610a88565b600b54604051600554604060020a909204600160a060020a03169160009182818181858883f1505050600555505b600b8054600054600160a060020a0316604060020a027bffffffffffffffffffffffffffffffffffffffff000000000000000019909116179055670de0b6b3a7640000600455565b610002565b42620d2f006003600050540111158015610ca457506064600460005054609602048110155b15610c7a57600b54600854604051604060020a909204600160a060020a0316916000919082818181858883f1505060068054606460058702040190555050506004819055600b80547bffffffffffffffffffffffffffffffffffffffff00000000000000001916604060020a3302179055426003556104778161094e565b61047760075467016345785d8a0000901061047957600754600154604051600160a060020a03919091169160009182818181858883f15050506007555056c65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8",
    address: "0x6bf68d0470d58001ca16f7b86f8c21d075f18bda",
    generated_with: "2.0.6",
    contract_name: "GameOfThrones"
  };

  function Contract() {
    if (Contract.Pudding == null) {
      throw new Error("GameOfThrones error: Please call load() first before creating new instance of this contract.");
    }

    Contract.Pudding.apply(this, arguments);
  };

  Contract.load = function(Pudding) {
    Contract.Pudding = Pudding;

    Pudding.whisk(contract_data, Contract);

    // Return itself for backwards compatibility.
    return Contract;
  }

  Contract.new = function() {
    if (Contract.Pudding == null) {
      throw new Error("GameOfThrones error: Please call load() first before calling new().");
    }

    return Contract.Pudding.new.apply(Contract, arguments);
  };

  Contract.at = function() {
    if (Contract.Pudding == null) {
      throw new Error("GameOfThrones error: lease call load() first before calling at().");
    }

    return Contract.Pudding.at.apply(Contract, arguments);
  };

  Contract.deployed = function() {
    if (Contract.Pudding == null) {
      throw new Error("GameOfThrones error: Please call load() first before calling deployed().");
    }

    return Contract.Pudding.deployed.apply(Contract, arguments);
  };

  if (typeof module != "undefined" && typeof module.exports != "undefined") {
    module.exports = Contract;
  } else {
    // There will only be one version of Pudding in the browser,
    // and we can use that.
    window.GameOfThrones = Contract;
  }

})();
