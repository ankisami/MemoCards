import 'package:flutter/material.dart';
import 'package:memocards/constants.dart';

class OptionDecksDetails extends StatelessWidget {
  final String title;
  final IconData iconData;
  final Color color;
  final Function press;

  const OptionDecksDetails({
    Key key,
    this.title,
    this.iconData,
    this.color,
    this.press,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(builder: (context, constraint) {
      return ClipRRect(
        borderRadius: BorderRadius.circular(13),
        child: Container(
          //margin: EdgeInsets.symmetric(vertical: 5, horizontal: 5),
          width: constraint.maxWidth / 2 -
              10, //constraint.maxWidth provide us the available width for this widget
          height: constraint.maxWidth / 3,
          //padding: EdgeInsets.all(16),
          decoration: BoxDecoration(color: color, boxShadow: [
            BoxShadow(
              offset: Offset(0, 17),
              blurRadius: 23,
              spreadRadius: -13,
              color: kShadowColor,
            )
          ]),
          child: Material(
            color: Colors.transparent,
            child: InkWell(
              onTap: press,
              child: Padding(
                padding: const EdgeInsets.all(16.0),
                child: Row(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: <Widget>[
                    Container(
                      height: 42,
                      width: 43,
                      decoration: BoxDecoration(
                        color: kBlueColor,
                        shape: BoxShape.circle,
                      ),
                      child: Icon(
                        iconData,
                        color: Colors.white,
                      ),
                    ),
                    SizedBox(width: 15),
                    Text(
                      title,
                      style: Theme.of(context).textTheme.subtitle,
                    ),
                  ],
                ),
              ),
            ),
          ),
        ),
      );
    });
  }
}
