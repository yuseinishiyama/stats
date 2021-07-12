//
//  StatsInputApp.swift
//  StatsInput
//
//  Created by Yusei Nishiyama on 11.07.21.
//

import SwiftUI

@main
struct StatsInputApp: App {
    @UIApplicationDelegateAdaptor private var appDelegate: AppDelegate
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}

class AppDelegate: NSObject, UIApplicationDelegate {
    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey : Any]? = nil) -> Bool {
        start()
        return true
    }    
}
