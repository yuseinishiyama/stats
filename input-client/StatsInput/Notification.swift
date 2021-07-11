//
//  Notification.swift
//  StatsInput
//
//  Created by Yusei Nishiyama on 12.07.21.
//

import Foundation
import NotificationCenter

func start() {
    UNUserNotificationCenter.current()
        .requestAuthorization(options: [.alert, .sound, .badge]) { granted, _  in
            UNUserNotificationCenter.current().removeAllPendingNotificationRequests()
            scheduleNotification()
        }
}

func scheduleNotification() {
    let content = UNMutableNotificationContent()
    content.body = "Rate your current mood"

    var date = DateComponents()
    date.hour = 0
    let trigger = UNCalendarNotificationTrigger(dateMatching: date, repeats: true)

    let request = UNNotificationRequest(
        identifier: UUID().uuidString,
        content: content,
        trigger: trigger)

    UNUserNotificationCenter.current().add(request) { error in
        if let error = error {
            print(error)
        }
    }
}
