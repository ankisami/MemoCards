import 'package:flutter/material.dart';
import 'package:memocards/constants.dart';

class FlashCardsDetails extends StatefulWidget {
  @override
  _FlashCardsDetailsState createState() => _FlashCardsDetailsState();
}

class _FlashCardsDetailsState extends State<FlashCardsDetails> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: kActiveIconColor,
        title: Text(
          "Anglais",
          style: Theme.of(context)
              .textTheme
              .display1
              .copyWith(fontWeight: FontWeight.w600),
        ),
      ),
      body: SafeArea(
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 10),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: <Widget>[
              Expanded(
                flex: 4,
                child: Padding(
                  padding: EdgeInsets.all(10.0),
                  child: Container(
                    height: 100,
                    decoration: BoxDecoration(
                        borderRadius: BorderRadius.all(Radius.circular(20)),
                        color: kActiveIconColor),
                  ),
                ),
              ),
              Expanded(
                child: Row(
                  crossAxisAlignment: CrossAxisAlignment.stretch,
                  children: <Widget>[
                    //ButtonAnswerCard(),
                    ButtonFlashCard(
                        title: "Facile", color: Colors.greenAccent[400]),
                    ButtonFlashCard(title: "Moyen", color: Colors.yellow),
                    ButtonFlashCard(title: "Je ne sais pas", color: Colors.red),
                  ],
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}

class ButtonAnswerCard extends StatelessWidget {
  const ButtonAnswerCard({
    Key key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return new InkWell(
      onTap: () => print('hello'),
      child: new Container(
        //width: 100.0,
        height: 50.0,
        decoration: new BoxDecoration(
          color: Colors.blueAccent,
          border: new Border.all(color: Colors.white, width: 2.0),
          borderRadius: new BorderRadius.circular(10.0),
        ),
        child: new Center(
          child: new Text(
            'Click Me',
            style: new TextStyle(fontSize: 18.0, color: Colors.white),
          ),
        ),
      ),
    );
  }
}

class ButtonFlashCard extends StatelessWidget {
  final String title;
  final Color color;
  const ButtonFlashCard({
    Key key,
    this.title,
    this.color,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Padding(
        padding: EdgeInsets.all(5.0),
        child: FlatButton(
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(18.0),
          ),
          textColor: Colors.white,
          color: color,
          child: Text(
            title,
            textAlign: TextAlign.center,
            style: TextStyle(
              color: Colors.black,
              fontSize: 20,
            ),
          ),
          onPressed: () {
            print('test');
          },
        ),
      ),
    );
  }
}
