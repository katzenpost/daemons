The Katzenpost Mix Network Daemons
==================================


Notes
-----

* The `daemons/vendor` directory contains all external dependencies
  that are required to build any part of Katzenpost, and can be
  reused if building things as libraries (Eg: symlinking
  `minclient/vendor` to `daemons/vendor` ).

* To build a new release see our release checklist document in our `docs`
  repository for the full release instructions.

* For a development work flow the Gopkg.toml may be edited to use the master
  branch of each Katzenpost repository.

License
-------

`AGPL <https://www.gnu.org/licenses/agpl-3.0.en.html>`_: see the LICENSE file
for details.  The external dependencies have their own licenses.


supported by
============

.. image:: https://katzenpost.mixnetworks.org/_static/images/eu-flag-tiny.jpg

This project has received funding from the European Union’s Horizon 2020
research and innovation programme under the Grant Agreement No 653497, Privacy
and Accountability in Networks via Optimized Randomized Mix-nets (Panoramix).
