import {describe, it, expect} from 'test/lib/common';

import {LogDisplayPlatformApp} from 'app/app';

describe('LogDisplayPlatformApp', () => {

  var app = new LogDisplayPlatformApp();

  it('can call inits', () => {
    expect(app).to.not.be(null);
  });
});


