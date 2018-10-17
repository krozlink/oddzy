/* eslint-disable */
import DateHelper from '../api/date-helper';

function past(minutes) {
	const past = new Date();
	past.setMinutes(past.getMinutes() - minutes);
	return past.getTime() / 1000;
}

function future(minutes) {
	const future = new Date();
	future.setMinutes(future.getMinutes() + minutes);
	return future.getTime() / 1000;
}

function now() {
	return new Date().getTime() / 1000;
}

function Today() {
    return {
      has_races: true,
      'date': DateHelper.formatDate(new Date()),
      'meetings': [{
        meeting_id: '3b493329-d5b9-418e-ae4b-a1c6132ef5e7',
        'name': 'Kyneton',
        country: 'Australia',
        'race_type': 'horse-racing',
        scheduled_start: past(30),
        'race_ids': ['9baa715e-d57e-4ee5-885c-f23b09a7ac30', 'f235ff87-a208-4c84-83d7-9809c11a01f5'],
        last_update: now(),
      },],
      'races': [{
        'race_id': '9baa715e-d57e-4ee5-885c-f23b09a7ac30',
        name: 'bet365 SV 3YO Fillies Maiden Plate',
        'number': 1,
        'status': 'OPEN',
        results: '',
        scheduled_start: past(30),
        last_update: now(),
      }, {
        'race_id': 'f235ff87-a208-4c84-83d7-9809c11a01f5',
        name: 'Kyneton Electrics Maiden Plate',
        'number': 2,
        'status': 'OPEN',
        'results': '',
        scheduled_start: future(30),
        'last_update': now(),
      }],
    };
}

function Tomorrow() {
  var d = new Date();
  d.setDate(d.getDate() + 1);

  return {
      "has_races": true,
      "date": DateHelper.formatDate(d),
      "meetings": [{
              "meeting_id": "d96b785a-66a9-4dec-ab30-adc847d9274b",
              "name": "Ballarat",
              "country": "Australia",
              "race_type": "harness",
              "scheduled_start": future(1000),
              "race_ids": ["ac0f6eef-bd19-4e29-a378-f3b62b357e48", "7e716650-adce-49cd-963b-b6640ed61d1e"],
              "last_update": 1539526574
          }, {
              "meeting_id": "006c70dd-8fa3-4c05-9d73-63b42a5be4e2",
              "name": "Caulfield",
              "country": "Australia",
              "race_type": "horse-racing",
              "scheduled_start":  future(1000),
              "race_ids": ["d2d24c4d-28dd-4afb-8df5-bae785aa8f10", "88e57ea5-9bfd-45d6-8a0f-99dbd02f725c"],
              "last_update": 1539682513
          }
      ],
      "races": [{
              "race_id": "ac0f6eef-bd19-4e29-a378-f3b62b357e48",
              "name": "R1 Dnr Logistics Pace",
              "number": 1,
              "status": "OPEN",
              "results": "",
              "scheduled_start": future(1490),
              "last_update": now()
          }, {
              "race_id": "7e716650-adce-49cd-963b-b6640ed61d1e",
              "name": "R2 Black Label Tattoo Collective Trot",
              "number": 2,
              "status": "OPEN",
              "results": "",
              "scheduled_start": future(1520),
              "last_update": now()
          }, {
              "race_id": "d2d24c4d-28dd-4afb-8df5-bae785aa8f10",
              "name": "Ladbrokes Multiverse Hcp",
              "number": 1,
              "status": "OPEN",
              "results": "",
              "scheduled_start": future(1310),
              "last_update": now()
          }, {
              "race_id": "88e57ea5-9bfd-45d6-8a0f-99dbd02f725c",
              "name": "DrinkWise Plate",
              "number": 2,
              "status": "OPEN",
              "results": "",
              "scheduled_start": future(1370),
              "last_update": now()
          }
      ]
  }    
}

function Overmorrow() {
  var d = new Date();
  d.setDate(d.getDate() + 2);

  return {
      "has_races": true,
      "date": DateHelper.formatDate(d),
      "meetings": [{
              "meeting_id": "5b1f3593-5426-4c54-802f-5a9b7f9c0fab",
              "name": "Happy Valley",
              "country": "Hong Kong",
              "race_type": "horse-racing",
              "scheduled_start": future(2880),
              "race_ids": ["ae0de897-8bda-4434-962f-2db1a08d0ef7", "c29ba64f-422e-4597-991d-c1aa6c2a7856", "e04599e2-8167-482a-a0e8-0a89a1150e79"],
              "last_update": now()
          }
      ],
      "races": [{
              "race_id": "ae0de897-8bda-4434-962f-2db1a08d0ef7",
              "name": "Aster Hcp (C4)",
              "number": 1,
              "status": "OPEN",
              "results": "",
              "scheduled_start": future(2880),
              "last_update": now()
          }, {
              "race_id": "c29ba64f-422e-4597-991d-c1aa6c2a7856",
              "name": "Aster Hcp (C4)",
              "number": 2,
              "status": "OPEN",
              "results": "",
              "scheduled_start": future(2910),
              "last_update": now()
          }, {
              "race_id": "e04599e2-8167-482a-a0e8-0a89a1150e79",
              "name": "Dandelion Hcp (C4)",
              "number": 3,
              "status": "OPEN",
              "results": "",
              "scheduled_start": future(2940),
              "last_update": now()
          }
      ]
  }
}

const races = {
	"9baa715e-d57e-4ee5-885c-f23b09a7ac30": {"race_id":"9baa715e-d57e-4ee5-885c-f23b09a7ac30","name":"bet365 SV 3YO Fillies Maiden Plate","number":1,"scheduled_start":1539657000,"results":"3,4,8","last_update":1539682513,"status":"CLOSED","meeting":{"meeting_id":"3b493329-d5b9-418e-ae4b-a1c6132ef5e7","name":"Kyneton","country":"Australia","race_type":"horse-racing","race_ids":["9baa715e-d57e-4ee5-885c-f23b09a7ac30","f235ff87-a208-4c84-83d7-9809c11a01f5","838334bf-81f1-4202-bffc-4969d319aaaf","4a091d92-77b6-46be-aa20-38c2a95c6037","ff04cc7a-93a7-4f5b-9031-dd942e903441","75f35e58-aaae-4f58-8017-d7b36b9f81ef","d4fa75b1-0ea0-438a-869f-683b169302e7","cce2a0fc-5855-4008-b692-8d300e858ec4"]},"selections":[{"selection_id":"f51fae0a-91b6-41d5-a8a3-7db046ccb476","name":"Extreme Pride","barrier":8,"number":1,"jockey":"John Keating","is_scratched":false,"weight":"57.0kg","jockey_weight":"55.5kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/666559_extreme_pride.png"},{"selection_id":"581ab714-7ffa-43ec-ac1e-1289dbdf90fc","name":"Infinite Spirit","barrier":10,"number":2,"jockey":"Toby Atkinson","is_scratched":false,"weight":"57.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/607525_infinite_spirit.png"},{"selection_id":"b000e6ce-0c9d-4d05-b5d8-d34fc0012320","name":"Jentico","barrier":9,"number":3,"jockey":"Stephanie Thornton","is_scratched":false,"weight":"57.0kg","jockey_weight":"51.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/666567_jentico.png"},{"selection_id":"6416175e-dbfd-4a97-9f8f-b9ed6e1516de","name":"Lady Solly","barrier":3,"number":4,"jockey":"Jason Baldock","is_scratched":false,"weight":"57.0kg","jockey_weight":"55.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/666568_lady_solly.png"},{"selection_id":"bb73b845-dc87-476d-ad3b-8111aae84e3d","name":"Lomachenko","barrier":11,"number":5,"jockey":"Joe Bowditch","is_scratched":false,"weight":"57.0kg","jockey_weight":"54.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/623660_lomachenko.png"},{"selection_id":"9ac473a3-16a7-47d0-a3cb-8e94f81c961a","name":"Lynn's Dream","barrier":4,"number":6,"jockey":"Brian Higgins","is_scratched":false,"weight":"57.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/624070_lynns_dream.png"},{"selection_id":"7a121b7b-f592-4eaa-9fc0-349d912728ee","name":"Niva","barrier":5,"number":7,"jockey":"Michael Dee","is_scratched":false,"weight":"57.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/654416_niva.png"},{"selection_id":"20ea3900-a81d-4d36-9120-74c82020f46c","name":"Noteably","barrier":2,"number":8,"jockey":"Ben Thompson","is_scratched":false,"weight":"57.0kg","jockey_weight":"48.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/666569_noteably.png"},{"selection_id":"c34c24e7-0cba-4e38-8e21-5f5fd06380bf","name":"Oggalution","barrier":1,"number":9,"jockey":"Michael Poy","is_scratched":false,"weight":"57.0kg","jockey_weight":"50.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/618631_oggalution.png"},{"selection_id":"6b15842e-c66d-44a5-b959-ba6d823b1dc8","name":"Sugar With That","barrier":7,"number":10,"jockey":"Ethan Brown","is_scratched":false,"weight":"57.0kg","jockey_weight":"56.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/618209_sugar_with_that.png"},{"selection_id":"a4073a23-b965-4882-b4ec-7b6623e3580d","name":"Tsushima","barrier":6,"number":11,"jockey":"Chris Symons","is_scratched":false,"weight":"57.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/666570_tsushima.png"}]},
	"f235ff87-a208-4c84-83d7-9809c11a01f5": {"race_id":"f235ff87-a208-4c84-83d7-9809c11a01f5","name":"Kyneton Electrics Maiden Plate","number":2,"scheduled_start":1539658800,"results":"","last_update":1539682513,"status":"OPEN","meeting":{"meeting_id":"3b493329-d5b9-418e-ae4b-a1c6132ef5e7","name":"Kyneton","country":"Australia","race_type":"horse-racing","race_ids":["9baa715e-d57e-4ee5-885c-f23b09a7ac30","f235ff87-a208-4c84-83d7-9809c11a01f5","838334bf-81f1-4202-bffc-4969d319aaaf","4a091d92-77b6-46be-aa20-38c2a95c6037","ff04cc7a-93a7-4f5b-9031-dd942e903441","75f35e58-aaae-4f58-8017-d7b36b9f81ef","d4fa75b1-0ea0-438a-869f-683b169302e7","cce2a0fc-5855-4008-b692-8d300e858ec4"]},"selections":[{"selection_id":"24eb6eb7-9818-4108-bb5d-239d4b0d97cd","name":"Generalmaintenance","barrier":13,"number":1,"jockey":"Dylan Dunn","is_scratched":false,"weight":"58.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/620605_generalmaintenance.png"},{"selection_id":"518145e4-387e-4508-83d5-a1fa2087f80f","name":"The Endorser","barrier":17,"number":2,"jockey":"Brad Rawiller","is_scratched":false,"weight":"58.5kg","jockey_weight":"54.5kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/575516_the_endorser.png"},{"selection_id":"82f43352-5ba4-496f-83b4-ab7f46f2dc3f","name":"Costalot","barrier":1,"number":3,"jockey":"Toby Atkinson","is_scratched":false,"weight":"56.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/528669_costalot.png"},{"selection_id":"c6652918-9142-4d56-ae8d-43e9ad7567d2","name":"Diva Peron","barrier":2,"number":4,"jockey":"Michael Dee","is_scratched":false,"weight":"56.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/568096_diva_peron.png"},{"selection_id":"7766e53c-6ad4-4512-9e41-0e3f8a2e8222","name":"Endless Rose","barrier":6,"number":5,"jockey":"Joe Bowditch","is_scratched":false,"weight":"56.5kg","jockey_weight":"54.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/595769_endless_rose.png"},{"selection_id":"f0a57ad5-b49b-4291-9210-7fc55ede501f","name":"Jinda","barrier":8,"number":6,"jockey":"Teodore Nugent","is_scratched":false,"weight":"56.5kg","jockey_weight":"49.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/605939_jinda.png"},{"selection_id":"1dda875e-32e0-4aa5-a107-c4fee37311c7","name":"Grand Symphony","barrier":10,"number":7,"jockey":"Brian Park","is_scratched":false,"weight":"56.0kg","jockey_weight":"54.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/604624_grand_symphony.png"}]},
	"ac0f6eef-bd19-4e29-a378-f3b62b357e48": {"race_id":"ac0f6eef-bd19-4e29-a378-f3b62b357e48","name":"R1 Dnr Logistics Pace","number":1,"scheduled_start":1539760980,"results":"","last_update":1539682787,"status":"OPEN","meeting":{"meeting_id":"d96b785a-66a9-4dec-ab30-adc847d9274b","name":"Ballarat","country":"Australia","race_type":"harness","race_ids":["ac0f6eef-bd19-4e29-a378-f3b62b357e48","7e716650-adce-49cd-963b-b6640ed61d1e","76cd3898-cae4-4663-b5ea-3ac3552e8544","52e25dc8-3eff-487c-a128-704ce05b8047","ff5adbb7-3904-46d9-b6a2-2a79f4432e16","ed345a7d-6313-4058-9cc0-acef131a1cf9","86a1b3a9-7ed2-4806-bed6-bf1479720063"]},"selections":[{"selection_id":"c3b0e7ca-5478-4d94-b44c-eef535f07ae8","name":"Oscar Bravo","barrier":0,"number":1,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"3d1a57c5-0860-4ae1-a5a3-4199541e49f9","name":"Blissfull Penny","barrier":0,"number":2,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"978681b1-1feb-4adf-bd5c-a069bdbf71ba","name":"Saviour Clare","barrier":0,"number":3,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"77379014-8359-468b-8dcc-84402c39fbd1","name":"Ta Failte Rote","barrier":0,"number":4,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"04c5ae4f-3fde-43f8-9dd2-e44d47b8eb27","name":"Sumarian Artist","barrier":0,"number":5,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"bd63e0e6-bed4-4f53-9eaa-c484502a0a22","name":"Be Major Threat","barrier":0,"number":6,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"30935a5c-54ea-4308-ab75-a6db79017236","name":"Terror Rising","barrier":0,"number":7,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"b6758a8a-efd7-4d9c-8ffd-b66df0a301f4","name":"Cocosfella","barrier":0,"number":8,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""}]},
	"7e716650-adce-49cd-963b-b6640ed61d1e": {"race_id":"7e716650-adce-49cd-963b-b6640ed61d1e","name":"R2 Black Label Tattoo Collective Trot","number":2,"scheduled_start":1539763200,"results":"","last_update":1539682788,"status":"OPEN","meeting":{"meeting_id":"d96b785a-66a9-4dec-ab30-adc847d9274b","name":"Ballarat","country":"Australia","race_type":"harness","race_ids":["ac0f6eef-bd19-4e29-a378-f3b62b357e48","7e716650-adce-49cd-963b-b6640ed61d1e","76cd3898-cae4-4663-b5ea-3ac3552e8544","52e25dc8-3eff-487c-a128-704ce05b8047","ff5adbb7-3904-46d9-b6a2-2a79f4432e16","ed345a7d-6313-4058-9cc0-acef131a1cf9","86a1b3a9-7ed2-4806-bed6-bf1479720063"]},"selections":[{"selection_id":"f86a5803-bad3-40d0-a77a-f758d2079aba","name":"Spiders Devil","barrier":0,"number":1,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"acad510b-8430-47dd-85f6-a30cf91ff3be","name":"Mel Durant","barrier":0,"number":2,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"1b8dd3a7-0570-4f59-a196-ed6008469a46","name":"Double Dot","barrier":0,"number":3,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"e682bd51-7a31-4044-b42d-4cd7d09371b7","name":"Penelope","barrier":0,"number":4,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"eca3376f-8849-4b14-adac-82edc81e806e","name":"Enticing Smile","barrier":0,"number":5,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"eff4dda3-abd4-465d-88f2-7b6c1136d2f0","name":"Coco Couture","barrier":0,"number":6,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"bf4a40f8-d5b2-42c7-a154-c8253f8c4eed","name":"Catchya Freshy","barrier":0,"number":7,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"b92dd889-9a39-4b70-8127-e8f2b0182433","name":"Kilkee","barrier":0,"number":8,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""},{"selection_id":"dc1be8f0-4588-4614-8e39-706cb038c496","name":"Peel My Grapes","barrier":0,"number":9,"jockey":"","is_scratched":false,"weight":"","jockey_weight":"","image_url":""}]},
	"d2d24c4d-28dd-4afb-8df5-bae785aa8f10": {"race_id":"d2d24c4d-28dd-4afb-8df5-bae785aa8f10","name":"Ladbrokes Multiverse Hcp","number":1,"scheduled_start":1539743400,"results":"","last_update":1539682957,"status":"OPEN","meeting":{"meeting_id":"006c70dd-8fa3-4c05-9d73-63b42a5be4e2","name":"Caulfield","country":"Australia","race_type":"horse-racing","race_ids":["d2d24c4d-28dd-4afb-8df5-bae785aa8f10","88e57ea5-9bfd-45d6-8a0f-99dbd02f725c","8df3f283-b3b6-4f5e-8680-29dfb7837c5d","d9266f20-3c46-455f-b41e-2de76fea5bfa","f5fcac86-bd01-437f-82b2-bcee26493561","97fa4032-e530-4c43-b2e5-068e5f8b4d63","ecdfdd27-8847-4338-881a-6756a952fd3c","fe91612b-0363-4b25-85de-a69ec24094a5"]},"selections":[{"selection_id":"515f0be7-fd13-4184-9585-4b90096f3e13","name":"Al Haram","barrier":2,"number":1,"jockey":"Regan Bayliss","is_scratched":false,"weight":"59.0kg","jockey_weight":"54.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/594252_al_haram.png"},{"selection_id":"0844ca55-9101-4b40-9c0f-558b4a36645c","name":"Valac","barrier":3,"number":2,"jockey":"James McDonald","is_scratched":false,"weight":"59.0kg","jockey_weight":"52.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/575511_valac.png"},{"selection_id":"10f293a0-877a-4b52-a812-585c9958dcb5","name":"Royal Music","barrier":4,"number":3,"jockey":"Craig Williams","is_scratched":false,"weight":"57.5kg","jockey_weight":"51.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/629216_royal_music.png"},{"selection_id":"fa3627fe-368b-4b8e-bd04-47aaba8b1b86","name":"Azuro","barrier":7,"number":4,"jockey":"Ben Allen","is_scratched":false,"weight":"56.0kg","jockey_weight":"50.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/629200_azuro.png"},{"selection_id":"a1c94ba9-d6a6-47df-9c7b-8150843bbf2f","name":"Steel Prince","barrier":6,"number":5,"jockey":"Dean Yendall","is_scratched":false,"weight":"55.5kg","jockey_weight":"50.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/535603_steel_prince.png"},{"selection_id":"8541925f-afb6-49f5-b3a9-fe76415c6102","name":"Opposition","barrier":5,"number":6,"jockey":"Kerrin McEvoy","is_scratched":false,"weight":"55.0kg","jockey_weight":"49.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/402208_opposition.png"},{"selection_id":"d39b2549-24c5-44d8-8b6e-22404aac68a4","name":"Fastlane to Heaven","barrier":1,"number":7,"jockey":"Dean Holland","is_scratched":false,"weight":"54.0kg","jockey_weight":"47.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/497868_fastlane_to_heaven.png"}]},
	"88e57ea5-9bfd-45d6-8a0f-99dbd02f725c": {"race_id":"88e57ea5-9bfd-45d6-8a0f-99dbd02f725c","name":"DrinkWise Plate","number":2,"scheduled_start":1539745200,"results":"","last_update":1539682958,"status":"OPEN","meeting":{"meeting_id":"006c70dd-8fa3-4c05-9d73-63b42a5be4e2","name":"Caulfield","country":"Australia","race_type":"horse-racing","race_ids":["d2d24c4d-28dd-4afb-8df5-bae785aa8f10","88e57ea5-9bfd-45d6-8a0f-99dbd02f725c","8df3f283-b3b6-4f5e-8680-29dfb7837c5d","d9266f20-3c46-455f-b41e-2de76fea5bfa","f5fcac86-bd01-437f-82b2-bcee26493561","97fa4032-e530-4c43-b2e5-068e5f8b4d63","ecdfdd27-8847-4338-881a-6756a952fd3c","fe91612b-0363-4b25-85de-a69ec24094a5"]},"selections":[{"selection_id":"51dde3f8-381b-4229-94cf-0f0356277280","name":"Real Success","barrier":3,"number":1,"jockey":"Ben E Thompson","is_scratched":false,"weight":"59.0kg","jockey_weight":"48.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/623883_real_success.png"},{"selection_id":"24955f1e-fa06-441b-b078-edd0364aa580","name":"Cloak","barrier":8,"number":2,"jockey":"James McDonald","is_scratched":false,"weight":"58.5kg","jockey_weight":"52.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/607614_cloak.png"},{"selection_id":"b7207ebf-a7a4-4f93-9d50-36a9dba0d481","name":"Silent Explorer","barrier":10,"number":3,"jockey":"Craig Williams","is_scratched":false,"weight":"58.5kg","jockey_weight":"51.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/622881_silent_explorer.png"},{"selection_id":"a86b3a18-0aa9-4d1c-aa1f-5caa1e550cc0","name":"The Founder","barrier":2,"number":5,"jockey":"Jordan Childs","is_scratched":false,"weight":"57.5kg","jockey_weight":"55.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/612043_the_founder.png"},{"selection_id":"a518437e-1b62-42e1-8b21-70106e635ad8","name":"A Fighting Fury","barrier":5,"number":6,"jockey":"Kerrin McEvoy","is_scratched":false,"weight":"56.5kg","jockey_weight":"49.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/611686_a_fighting_fury.png"},{"selection_id":"a14176f4-11c3-4156-acb4-77bca1cfcd13","name":"Junipal","barrier":4,"number":7,"jockey":"John Allen","is_scratched":false,"weight":"56.5kg","jockey_weight":"56.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/647452_junipal.png"},{"selection_id":"66d34710-0f67-4bb0-95cd-93ec6ea60c04","name":"Gheedaa","barrier":9,"number":8,"jockey":"Regan Bayliss","is_scratched":false,"weight":"56.0kg","jockey_weight":"54.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/601226_gheedaa.png"},{"selection_id":"5ca658d8-12ee-4f59-aed3-cdbbcd01407e","name":"Mr Grizzle","barrier":1,"number":9,"jockey":"Damian Lane","is_scratched":false,"weight":"56.0kg","jockey_weight":"53.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/612619_mr_grizzle.png"},{"selection_id":"3d743f84-4f20-45be-b577-a0a6050114f5","name":"Predecessor","barrier":12,"number":10,"jockey":"Dwayne Dunn","is_scratched":false,"weight":"56.0kg","jockey_weight":"53.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/645070_predecessor.png"},{"selection_id":"07a52308-5485-4e90-84ca-8ae78159512d","name":"Syd's Coin","barrier":6,"number":11,"jockey":"Linda Meech","is_scratched":false,"weight":"55.0kg","jockey_weight":"54.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/605934_syds_coin.png"},{"selection_id":"a66cfcc6-75ee-4fdd-b219-e7669979b366","name":"Titan Blinders","barrier":7,"number":12,"jockey":"Damien Oliver","is_scratched":false,"weight":"55.0kg","jockey_weight":"54.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/605958_titan_blinders.png"},{"selection_id":"cd7b2cc7-3c11-488a-a393-2d414e71e645","name":"Yulong Meteor","barrier":11,"number":13,"jockey":"Michael Walker","is_scratched":false,"weight":"54.5kg","jockey_weight":"52.0kg","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/606531_yulong_meteor.png"}]},
	"ae0de897-8bda-4434-962f-2db1a08d0ef7": {"race_id":"ae0de897-8bda-4434-962f-2db1a08d0ef7","name":"Aster Hcp (C4)","number":1,"scheduled_start":1539861300,"results":"","last_update":1539677057,"status":"OPEN","meeting":{"meeting_id":"5b1f3593-5426-4c54-802f-5a9b7f9c0fab","name":"Happy Valley","country":"Hong Kong","race_type":"horse-racing","race_ids":["ae0de897-8bda-4434-962f-2db1a08d0ef7","c29ba64f-422e-4597-991d-c1aa6c2a7856","e04599e2-8167-482a-a0e8-0a89a1150e79","7f873e00-900f-4319-9007-4336920c149a","316e833b-1c7e-4796-b1b1-501388fc4e7a","4b1ed035-d703-44b1-96aa-67230c8fbb45","d99676ab-77cc-4aca-b10d-d26a07c50b06","3342102b-914d-48b7-921a-ec24ab1e40b0"]},"selections":[{"selection_id":"fb5efe5d-527c-411a-ba42-e674f21e1880","name":"Good Companion","barrier":7,"number":1,"jockey":"Matthew Chadwick","is_scratched":false,"weight":"59.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/505269_good_companion.png"},{"selection_id":"b4d012e8-ddd6-418a-93a5-5aae47231f86","name":"Very Rich Man","barrier":1,"number":2,"jockey":"Matthew Poon","is_scratched":false,"weight":"59.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/465944_very_rich_man.png"},{"selection_id":"6d122765-7dc3-40e6-b9a7-adf0c39a6f42","name":"Pakistan Baby","barrier":4,"number":3,"jockey":"Karis Teetan","is_scratched":false,"weight":"58.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/301426_pakistan_baby.png"},{"selection_id":"74871528-9ae6-485f-9836-e53914d2a468","name":"The Createth","barrier":9,"number":4,"jockey":"Neil Callan","is_scratched":false,"weight":"58.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/458523_the_createth.png"},{"selection_id":"b8e607db-9ed0-4137-b94b-a41d54dbe469","name":"Spring Win","barrier":10,"number":5,"jockey":"Dylan Mo","is_scratched":false,"weight":"57.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/424442_spring_win.png"},{"selection_id":"57ad6e8d-d0e2-4973-be0a-3ca39d9f61a6","name":"Gracious Ryder","barrier":3,"number":6,"jockey":"Douglas Whyte","is_scratched":false,"weight":"56.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/526153_gracious_ryder.png"},{"selection_id":"3aaf2d48-3a14-41a9-b95e-0bb0a6dcbf36","name":"Prawn Yeah Yeah","barrier":6,"number":7,"jockey":"Callan Murray","is_scratched":false,"weight":"56.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/667151_prawn_yeah_yeah.png"},{"selection_id":"dd97035d-aa9f-4ddd-8665-35b61e7f0c4e","name":"London Master","barrier":8,"number":8,"jockey":"Umberto Rispoli","is_scratched":false,"weight":"56.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/390969_london_master.png"},{"selection_id":"de4869bb-e712-4f63-91b0-33e4953be6ee","name":"Junzi","barrier":5,"number":9,"jockey":"Alberto Sanna","is_scratched":false,"weight":"54.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/412416_junzi.png"},{"selection_id":"c24ca9f0-1253-4320-bbbc-d599c9cc594a","name":"Jumbo Bus","barrier":2,"number":10,"jockey":"Ben So","is_scratched":false,"weight":"53.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/593879_jumbo_bus.png"}]},
	"c29ba64f-422e-4597-991d-c1aa6c2a7856": {"race_id":"c29ba64f-422e-4597-991d-c1aa6c2a7856","name":"Aster Hcp (C4)","number":2,"scheduled_start":1539863100,"results":"","last_update":1539677059,"status":"OPEN","meeting":{"meeting_id":"5b1f3593-5426-4c54-802f-5a9b7f9c0fab","name":"Happy Valley","country":"Hong Kong","race_type":"horse-racing","race_ids":["ae0de897-8bda-4434-962f-2db1a08d0ef7","c29ba64f-422e-4597-991d-c1aa6c2a7856","e04599e2-8167-482a-a0e8-0a89a1150e79","7f873e00-900f-4319-9007-4336920c149a","316e833b-1c7e-4796-b1b1-501388fc4e7a","4b1ed035-d703-44b1-96aa-67230c8fbb45","d99676ab-77cc-4aca-b10d-d26a07c50b06","3342102b-914d-48b7-921a-ec24ab1e40b0"]},"selections":[{"selection_id":"c816a065-94c3-46f9-acc5-d59120b664ef","name":"Burst Away","barrier":9,"number":1,"jockey":"Douglas Whyte","is_scratched":false,"weight":"60.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/509078_burst_away.png"},{"selection_id":"e54a32dc-225e-4687-a8f9-ca320ebcfe78","name":"Fantastic Fabio","barrier":5,"number":2,"jockey":"Cash Wong","is_scratched":false,"weight":"57.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/495054_fantastic_fabio.png"},{"selection_id":"b9283fb0-a863-4290-bf3d-d77842ddc62a","name":"Big Bully","barrier":8,"number":3,"jockey":"Keith Yeung","is_scratched":false,"weight":"56.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/620421_big_bully.png"},{"selection_id":"a86a093d-9525-4417-a540-63bef6a6b5c0","name":"Young Empire","barrier":3,"number":4,"jockey":"Alberto Sanna","is_scratched":false,"weight":"56.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/429017_young_empire.png"},{"selection_id":"c65510b3-9b86-4404-80a0-1af95350925c","name":"Split Of A Second","barrier":1,"number":5,"jockey":"Matthew Poon","is_scratched":false,"weight":"55.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/601139_split_of_a_second.png"},{"selection_id":"40bc58e1-9b3a-4b13-afc2-74d252aba7c6","name":"Tom's Dragon","barrier":6,"number":6,"jockey":"Sam Clipperton","is_scratched":false,"weight":"55.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/618007_toms_dragon.png"},{"selection_id":"19b98d97-8a1d-42e8-8a24-88ea06f10c0e","name":"Love Chunghwa","barrier":4,"number":7,"jockey":"Umberto Rispoli","is_scratched":false,"weight":"55.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/414896_love_chunghwa.png"},{"selection_id":"d81c65cf-44e4-40dd-8581-8c6446d963e4","name":"Touchdown Striker","barrier":10,"number":8,"jockey":"Dylan Mo","is_scratched":false,"weight":"55.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/627225_touchdown_striker.png"},{"selection_id":"81faae3a-123b-4159-92de-59d75f17fd25","name":"Dominator","barrier":2,"number":9,"jockey":"Zac Purton","is_scratched":false,"weight":"54.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/623670_dominator.png"},{"selection_id":"2960880a-5c05-40d6-8f02-f10ff387da3c","name":"Manful Star","barrier":7,"number":10,"jockey":"Grant Van Niekerk","is_scratched":false,"weight":"53.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/443814_manful_star.png"}]},
	"e04599e2-8167-482a-a0e8-0a89a1150e79": {"race_id":"e04599e2-8167-482a-a0e8-0a89a1150e79","name":"Dandelion Hcp (C4)","number":3,"scheduled_start":1539864900,"results":"","last_update":1539677061,"status":"OPEN","meeting":{"meeting_id":"5b1f3593-5426-4c54-802f-5a9b7f9c0fab","name":"Happy Valley","country":"Hong Kong","race_type":"horse-racing","race_ids":["ae0de897-8bda-4434-962f-2db1a08d0ef7","c29ba64f-422e-4597-991d-c1aa6c2a7856","e04599e2-8167-482a-a0e8-0a89a1150e79","7f873e00-900f-4319-9007-4336920c149a","316e833b-1c7e-4796-b1b1-501388fc4e7a","4b1ed035-d703-44b1-96aa-67230c8fbb45","d99676ab-77cc-4aca-b10d-d26a07c50b06","3342102b-914d-48b7-921a-ec24ab1e40b0"]},"selections":[{"selection_id":"f0f00091-4b41-4c8f-9bef-395bdff07dc4","name":"Sichuan Boss","barrier":3,"number":1,"jockey":"Cash Wong","is_scratched":false,"weight":"60.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/522807_sichuan_boss.png"},{"selection_id":"10f8b83b-5f3e-485d-b385-1dc769c5b5f3","name":"Waldorf","barrier":2,"number":2,"jockey":"Alvin Ng","is_scratched":false,"weight":"59.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/493072_waldorf.png"},{"selection_id":"2dc4f3e2-2b07-4960-bd13-aa554177a4c3","name":"True Grit","barrier":1,"number":3,"jockey":"Sam Clipperton","is_scratched":false,"weight":"56.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/653545_true_grit.png"},{"selection_id":"749a3f40-aa5e-406d-a0d0-1c15ec78198c","name":"Captain Boss","barrier":8,"number":4,"jockey":"Zac Purton","is_scratched":false,"weight":"56.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/593405_captain_boss.png"},{"selection_id":"672c172b-e2e7-445b-9d5b-66185fa02828","name":"Golden Kid","barrier":9,"number":5,"jockey":"Alberto Sanna","is_scratched":false,"weight":"55.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/544778_golden_kid.png"},{"selection_id":"41aeb63d-db3d-4d9b-8a05-4451453c6f09","name":"Letsgofree","barrier":7,"number":6,"jockey":"Derek Leung","is_scratched":false,"weight":"54.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/409589_letsgofree.png"},{"selection_id":"25aa1988-4500-445b-a389-a1bd056178f3","name":"Horse Prosperous","barrier":5,"number":7,"jockey":"Vincent Ho","is_scratched":false,"weight":"53.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/589933_horse_prosperous.png"},{"selection_id":"b58925bc-cecd-4deb-8ab9-41e28ebda7be","name":"The Sylph","barrier":10,"number":8,"jockey":"Umberto Rispoli","is_scratched":false,"weight":"53.5kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/325438_the_sylph.png"},{"selection_id":"9f387673-42fc-4dda-adec-7d14f01f2b26","name":"Chaparral Star","barrier":6,"number":9,"jockey":"Karis Teetan","is_scratched":false,"weight":"52.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/557093_chaparral_star.png"},{"selection_id":"be4a1d36-dbe2-4779-9b68-85c0c46f107d","name":"Money Marshal","barrier":4,"number":10,"jockey":"Matthew Poon","is_scratched":false,"weight":"52.0kg","jockey_weight":"","image_url":"//dnu5embx6omws.cloudfront.net/silks/horse-racing/625884_money_marshal.png"}]},
}

function RaceSchedule(date) {
    var today = DateHelper.formatDate(DateHelper.todayDate());
    if(date === today) return Today();

    var tomorrow = DateHelper.formatDate(DateHelper.tomorrowDate());
    if(date === tomorrow) return Tomorrow();

    var overmorrow = DateHelper.formatDate(DateHelper.overmorrowDate());
    if(date === overmorrow) return Overmorrow();

    return {
        has_races: false,
        'date': date,
        'meetings': [],
        'races': [],
    };
}

function RaceCard(raceId) {
    return races[raceId];
}

export default {
  RaceSchedule,
  RaceCard,
};
