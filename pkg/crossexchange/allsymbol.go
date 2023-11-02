package crossexchange

// type SymbolFundingPriceDiffs map[string][]FundingPriceDiff

// func (exs *StrategyExecuter) GetCrossExchangeArbitrage() {
// 	//start := time.Now().UnixMilli()

// 	wg := new(sync.WaitGroup)
// 	wg.Add(len(exs.Exchanges))

// 	lock := new(sync.Mutex)
// 	result := domain.FundingPricesOfSymbol{}
// 	_ = result
// 	for _, ex := range exs.Exchanges {
// 		go func(ex any) {
// 			defer wg.Done()
// 			strat, ok := ex.(domain.GetFundingAndPricer)
// 			if !ok {
// 				log.Fatal(fmt.Errorf("type error: %v", ex))
// 			}

// 			info, err := (strat).GetFundingAndPrices()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			lock.Lock()
// 			for k, v := range info {
// 				result[k] = append(result[k], v...)
// 			}
// 			lock.Unlock()
// 		}(ex)
// 	}
// 	wg.Wait()

// 	results := SymbolFundingPriceDiffs{}
// 	for symbol, fundingPrices := range result {

// 		results[symbol] = calculateFundingPrices(fundingPrices)
// 	}
// 	fmt.Println(results["POLYXUSDT"])
// }
