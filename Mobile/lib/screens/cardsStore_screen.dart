import 'package:flutter/material.dart';
import 'package:memocards/widgets/bottom_nav_bar.dart';

class CardStoreScreen extends StatefulWidget {
  @override
  _CardStoreScreenState createState() => _CardStoreScreenState();
}

class _CardStoreScreenState extends State<CardStoreScreen>
    with SingleTickerProviderStateMixin {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: BottomNavBar(),
      body: Column(
        children: <Widget>[
          Padding(
            padding: const EdgeInsets.only(left: 20, right: 20, top: 50),
            child: Row(
              children: <Widget>[
                Icon(Icons.menu, color: Colors.black, size: 36),
                Spacer(),
                Icon(Icons.shopping_cart, color: Colors.black, size: 36),
              ],
            ),
          ),
          SizedBox(height: 35),
          Padding(
            padding: const EdgeInsets.only(left: 20.0, right: 20.0),
            child: Container(
              width: MediaQuery.of(context).size.width,
              height: 50,
              decoration: BoxDecoration(
                color: Colors.white,
                borderRadius: BorderRadius.circular(30),
                border: Border.all(color: Colors.grey, width: 0.5),
              ),
              child: Padding(
                padding: const EdgeInsets.only(left: 20.0, right: 20.0),
                child: Row(
                  children: <Widget>[
                    Text(
                      "Rechercher un deck ici",
                      style: TextStyle(
                          color: Colors.grey,
                          fontSize: 16,
                          fontFamily: 'OpenSans'),
                    ),
                    Spacer(),
                    Icon(Icons.search, color: Colors.black, size: 36),
                  ],
                ),
              ),
            ),
          ),
          SizedBox(height: 35),
          TabBar(
            indicatorColor: Colors.green,
            indicatorWeight: 3.0,
            labelColor: Colors.black,
            unselectedLabelColor: Colors.grey,
            isScrollable: true,
            tabs: <Widget>[
              Tab(
                child: Text('aaa', style: TextStyle(fontSize: 16)),
              ),
              Tab(
                child: Text('bbb', style: TextStyle(fontSize: 16)),
              ),
              Tab(
                child: Text('ccc', style: TextStyle(fontSize: 16)),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
