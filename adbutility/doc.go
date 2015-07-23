// Package adbutility provides useful methods to communicate with underlying
// adb tool. All device interactions are done via this package and this is the
// lowest level in the method call hierarchy.
//
// Following are the features provided by this package.
//    * Ability to communicate with adb endpoint on local machine
//    * Execute adb commands on local machine and get the output back
//    * Get list of attached devices serial numbers
//    * Wait for a set of serial numbers to be attached to the adb endpoint
//    * Wait for a specific number of devices to be attached to the adb endpoint
//    * Each adb call is logged for debugging purpose
//    * Each adb command is associated with a timeout slot, and the process is automatically killed after timeout.
package adbutility
