package controllers

import play.api._
import play.api.mvc._

import models._

object Application extends Controller {

    /**
     * Home
     *
     */
    def index = Action {
        Async {
            val response = Content.getContent
            response.map({ sheet =>
                Ok(views.html.index(Content.parse(sheet)))
            })
        }
    }

    /**
     * About
     *
     */
    def about = Action {
        Async {
            val response = Content.getContent
            response.map({ sheet =>
                Ok(views.html.about(Content.parse(sheet)))
            })
        }
    }

    /**
     * Solutions
     *
     */
    def solutions = Action {
        Async {
            val response = Content.getContent
            response.map({ sheet =>
                Ok(views.html.solutions(Content.parse(sheet)))
            })
        }
    }

    /**
     * Experience
     *
     */
    def experience = Action {
        Async {
            val response = Content.getContent
            response.map({ sheet =>
                Ok(views.html.experience(Content.parse(sheet)))
            })
        }
    }

    /**
     * Contact
     *
     */
    def contact = Action {
        Async {
            val response = Content.getContent
            response.map({ sheet =>
                Ok(views.html.contact(Content.parse(sheet)))
            })
        }
    }

    /**
     * Expo
     *
     */
    def expos(tag:String) = Action {
        Async {
            val response = Expos.getContent
            response.map({ sheet =>
                var expos = Expos.parse(sheet)
                if (tag != "Upcoming") expos = expos.filter(e => e.categorySlug == tag)
                val title = {
                    if (tag == "Upcoming") {
                        tag
                    } else if (expos.length > 0) {
                        expos(0).categoryName
                    } else {
                        "None found"
                    }
                }
                Ok(views.html.expos(title, expos))
            })
        }
    }
}
