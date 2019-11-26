# flexpipe

A library to ease pipeline data manipulations

This is a collection of types & functions enable easily writing pipelines to
filter and manipulate data leveraging goroutines and channels.

For example:

    flex.ReadCSV("input.csv").Filter(onlyCertainRecords).Map(convertSomeFields).OutputAsJSON()

    func onlyCertainRecords(r flex.Record) bool {...}

    func convertSomeFields(r flex.Record) flex.Record {...}

