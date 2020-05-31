import 'package:flutter/material.dart';
import 'package:memocards/widgets/bottom_nav_bar.dart';

class TodayScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: BottomNavBar(),
      body: Center(
        child: Text("Today page..."),
      ),
    );
  }
}
